//+build wireinject

package giro

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/wire"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"github.com/rerost/giro/domain/host"
	"github.com/rerost/giro/domain/message"
	"github.com/rerost/giro/domain/messagename"
	"github.com/rerost/giro/domain/service"
	hosts_pb "github.com/rerost/giro/pb/hosts"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type ReflectionAddr string
type RPCAddr string
type Metadata map[string]string

func NewServerReflectionConn(ctx context.Context, reflectionAddr ReflectionAddr) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, string(reflectionAddr), grpc.WithInsecure())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}

func NewServerReflectionClient(ctx context.Context, conn *grpc.ClientConn) (*grpcreflect.Client, error) {
	client := grpcreflect.NewClient(ctx, grpc_reflection_v1alpha.NewServerReflectionClient(conn))

	return client, nil
}

func ProviderReflectionAddr() ReflectionAddr {
	return ReflectionAddr(config.ReflectionServer)
}

func ProviderRPCAddr() RPCAddr {
	return RPCAddr(config.RpcServer)
}

func ProviderHostResolver(conn *grpc.ClientConn, rpcAddr RPCAddr) (host.HostResolver, error) {
	if rpcAddr != "" {
		return host.NewConstHostResolver(string(rpcAddr)), nil
	}

	client := hosts_pb.NewHostServiceClient(conn)

	return host.NewHostResolver(client), nil
}

var fromConfigSet = wire.NewSet(
	ProviderReflectionAddr,
	ProviderRPCAddr,
	ProviderHostResolver,
)

type LsCmd *cobra.Command

func ProviderLsCmd() LsCmd {
	cmd := &cobra.Command{
		Use:  "ls [service|method]",
		Args: cobra.MaximumNArgs(1),
		RunE: func(ccmd *cobra.Command, arg []string) error {
			ctx := ccmd.Context()
			serviceService, err := NewServiceService(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			var args []string
			if len(arg) == 1 {
				args = strings.Split(arg[0], "/")
			}

			switch len(args) {
			case 0:
				srvs, err := serviceService.Ls(ctx, nil, nil)
				if err != nil {
					return errors.WithStack(err)
				}
				for _, s := range srvs {
					fmt.Println(s.Name)
				}
			case 1:
				srvs, err := serviceService.Ls(ctx, &args[0], nil)
				if err != nil {
					return errors.WithStack(err)
				}
				for _, mn := range srvs[0].MethodNames {
					fmt.Println(mn)
				}
			case 2:
				srvs, err := serviceService.Ls(ctx, &args[0], &args[1])
				if err != nil {
					return errors.WithStack(err)
				}
				for _, mn := range srvs[0].MethodNames {
					fmt.Println(mn)
				}
			}

			return nil
		},
	}
	return cmd
}

type EmptyJSONCmd *cobra.Command

func ProviderEmptyJSONCmd() EmptyJSONCmd {
	cmd := &cobra.Command{
		Use:  "empty_json <message>",
		Args: cobra.ExactArgs(1),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()
			messageeService, err := NewMessageService(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			json, err := messageeService.EmptyJSON(ctx, messagename.MessageName(args[0]))
			if err != nil {
				return errors.WithStack(err)
			}

			fmt.Println(string(json))

			return nil
		},
	}

	return cmd
}

type ToJSONCmd *cobra.Command

func ProviderToJSONCmd() ToJSONCmd {
	cmd := &cobra.Command{
		Use:  "tojson <message> [message_body]",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()
			messageeService, err := NewMessageService(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			var body string
			if len(args) == 2 {
				body = args[1]
			} else {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return errors.WithStack(err)
				}
				body = string(b)
			}
			json, err := messageeService.ToJSON(ctx, messagename.MessageName(args[0]), message.Binary(body))
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Println(string(json))

			return nil
		},
	}
	return cmd
}

type ToBinaryCmd *cobra.Command

func ProviderToBinaryCmd() ToBinaryCmd {
	cmd := &cobra.Command{
		Use:  "tobinary <message> [message_body]",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()
			messageeService, err := NewMessageService(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			var body string
			if len(args) == 2 {
				body = args[1]
			} else {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return errors.WithStack(err)
				}
				body = string(b)
			}
			bin, err := messageeService.ToBinary(ctx, messagename.MessageName(args[0]), message.JSON(body))
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Println(string(bin))

			return nil
		},
	}

	return cmd
}

type CallCmd *cobra.Command

func ParseMetadata(m string) (map[string]string, error) {
	if m == "" {
		return nil, nil
	}
	ms := strings.Split(m, ":")
	if len(ms)%2 != 0 {
		return nil, errors.New("Expect key1:val1:key2:val2 format")
	}
	md := make(map[string]string, len(ms)/2)

	for i := 0; i < len(ms); i += 2 {
		md[ms[i]] = ms[i+1]
	}
	return md, nil
}

func ProviderCallCmd() CallCmd {
	metadata := ""
	cmd := &cobra.Command{
		Use:  "call <method> [message_body]",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()
			serviceService, err := NewServiceService(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			var body string
			if len(args) == 2 {
				body = args[1]
			} else {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return errors.WithStack(err)
				}
				body = string(b)
			}

			parsedMeataData, err := ParseMetadata(metadata)
			if err != nil {
				return errors.WithStack(err)
			}
			for k, v := range parsedMeataData {
				zap.L().Debug("received metadata", zap.String(k, v))
			}

			tmp := strings.Split(args[0], "/")
			svcName := tmp[0]
			methodName := tmp[1]
			bin, err := serviceService.Call(ctx, svcName, methodName, parsedMeataData, message.JSON(body))
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Println(string(bin))
			return nil
		},
	}

	cmd.Flags().StringVarP(&metadata, "metadata", "m", "", "metadata. e.g key1:val1:key2:val2")

	return cmd
}

type VersionCmd *cobra.Command
type Version string
type Revision string

func ProviderVersionCmd(version Version, revision Revision) (VersionCmd, error) {
	cmd := &cobra.Command{
		Use: "version",
		RunE: func(ccmd *cobra.Command, args []string) error {
			fmt.Printf("Version=%s, Revision=%s\n", version, revision)
			return nil
		},
	}

	return cmd, nil
}

type HostCmd *cobra.Command

func ProviderHostCmd() (HostCmd, error) {
	cmd := &cobra.Command{
		Use:  "host <service>",
		Args: cobra.ExactArgs(1),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()
			hostResolver, err := NewHostResolver(ctx, ccmd.Flags())
			if err != nil {
				return errors.WithStack(err)
			}

			host, err := hostResolver.Resolve(ctx, args[0])
			if err != nil {
				return errors.WithStack(err)
			}

			fmt.Println(host)
			return nil
		},
	}

	return cmd, nil
}

var config Config

func ProviderCmdRoot(lsCmd LsCmd, emptyJSONCmd EmptyJSONCmd, toJSONCmd ToJSONCmd, toBinaryCmd ToBinaryCmd, callCmd CallCmd, versionCmd VersionCmd, hostCmd HostCmd) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "giro",
		Short: "",
	}

	cmd.AddCommand(
		lsCmd,
		emptyJSONCmd,
		toJSONCmd,
		toBinaryCmd,
		callCmd,
		versionCmd,
		hostCmd,
	)

	cmd.PersistentFlags().StringP("reflection-server", "r", "localhost:5000", "")
	cmd.PersistentFlags().StringP("rpc-server", "", "", "")
	cmd.PersistentFlags().BoolP("verbose", "", false, "")
	cmd.PersistentFlags().BoolP("debug", "", false, "")

	cobra.OnInitialize(func() {
		flags := cmd.PersistentFlags()
		v := viper.New()
		v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		v.AutomaticEnv()
		v.BindPFlags(cmd.Flags())
		flags.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{
			UnknownFlags: true,
		}

		err := flags.Parse(os.Args[1:])
		if err != nil {
			panic(err)
		}

		var cfg Config
		err = v.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}

		config = cfg

		zcfg := zap.NewProductionConfig()
		if cfg.Debug {
			zcfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		}
		if cfg.Verbose {
			zcfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		}
		l, err := zcfg.Build()
		// TODO call l.Sync()

		zap.ReplaceGlobals(l)
	})

	return cmd, nil
}

func NewCmdRoot(ctx context.Context, version Version, revision Revision) (*cobra.Command, error) {
	wire.Build(
		ProviderCmdRoot,
		ProviderLsCmd,
		ProviderEmptyJSONCmd,
		ProviderToJSONCmd,
		ProviderToBinaryCmd,
		ProviderCallCmd,
		ProviderVersionCmd,
		ProviderHostCmd,
	)

	return nil, nil
}

var base = wire.NewSet(
	service.NewServiceService,
	message.NewMessageService,
	messagename.NewMessageNameResolver,
	NewServerReflectionClient,
	grpcreflectiface.NewClient,
	NewServerReflectionConn,
)

func NewServiceService(context.Context, *pflag.FlagSet) (service.ServiceService, error) {
	wire.Build(
		fromConfigSet,
		base,
	)

	return nil, nil
}

func NewMessageService(context.Context, *pflag.FlagSet) (message.MessageService, error) {
	wire.Build(
		fromConfigSet,
		base,
	)

	return nil, nil
}

func NewMessageNameResolver(context.Context, *pflag.FlagSet) (messagename.MessageNameResolver, error) {
	wire.Build(
		fromConfigSet,
		base,
	)

	return nil, nil
}

func NewHostResolver(context.Context, *pflag.FlagSet) (host.HostResolver, error) {
	wire.Build(
		fromConfigSet,
		base,
	)

	return nil, nil
}
