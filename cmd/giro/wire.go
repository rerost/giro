package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/wire"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/host"
	"github.com/rerost/giro/domain/message"
	"github.com/rerost/giro/domain/messagename"
	"github.com/rerost/giro/domain/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type ReflectionAddr string
type RPCAddr string

func NewServerReflectionClient(ctx context.Context, reflectionAddr ReflectionAddr) (*grpcreflect.Client, error) {
	conn, err := grpc.DialContext(ctx, string(reflectionAddr), grpc.WithInsecure())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := grpcreflect.NewClient(ctx, grpc_reflection_v1alpha.NewServerReflectionClient(conn))

	return client, nil
}

func ProvideReflectionAddr(cfg Config) ReflectionAddr {
	return ReflectionAddr(cfg.ReflectionServer)
}

func ProvideRPCAddr(cfg Config) RPCAddr {
	return RPCAddr(cfg.RpcServer)
}

func ProviderHostResolver(rpcAddr RPCAddr) (host.HostResolver, error) {
	if rpcAddr == "" {
		return nil, nil
	}

	return host.NewConstHostResolver(string(rpcAddr)), nil
}

var base = wire.NewSet(
	service.NewServiceService,
	message.NewMessageService,
	messagename.NewMessageNameResolver,
	NewServerReflectionClient,
	ProviderHostResolver,
	ProvideReflectionAddr,
	ProvideRPCAddr,
)

type LsCmd *cobra.Command

func ProviderLsCmd(serviceService service.ServiceService) LsCmd {
	cmd := &cobra.Command{
		Use:  "ls",
		Args: cobra.MaximumNArgs(2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()

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
				return errors.New("Unsupported")
			}

			return nil
		},
	}
	return cmd
}

type EmptyJSONCmd *cobra.Command

func ProviderEmptyJSONCmd(messageeService message.MessageService) EmptyJSONCmd {
	cmd := &cobra.Command{
		Use:  "empty_json",
		Args: cobra.ExactArgs(1),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()

			json, err := messageeService.EmptyJSON(ctx, messagename.MessageName(args[0]))
			if err != nil {
				return errors.WithStack(err)
			}

			fmt.Println(json)

			return nil
		},
	}

	return cmd
}

type ToJSONCmd *cobra.Command

func ProviderToJSONCmd(messageeService message.MessageService) ToJSONCmd {
	cmd := &cobra.Command{
		Use:  "tojson",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()

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
			fmt.Println(json)

			return nil
		},
	}
	return cmd
}

type ToBinaryCmd *cobra.Command

func ProviderToBinaryCmd(messageeService message.MessageService) ToBinaryCmd {
	cmd := &cobra.Command{
		Use:  "tobinary",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()

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
			fmt.Println(bin)

			return nil
		},
	}

	return cmd
}

type CallCmd *cobra.Command

func ProviderCallCmd(serviceService service.ServiceService) CallCmd {
	cmd := &cobra.Command{
		Use:  "call",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(ccmd *cobra.Command, args []string) error {
			ctx := ccmd.Context()

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

			tmp := strings.Split(args[0], "/")
			svcName := tmp[0]
			methodName := tmp[1]
			bin, err := serviceService.Call(ctx, svcName, methodName, nil, message.JSON(body))
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Println(bin)
			return nil
		},
	}

	return cmd
}

func ProviderCmdRoot(lsCmd LsCmd, emptyJSONCmd EmptyJSONCmd, toJSONCmd ToJSONCmd, toBinaryCmd ToBinaryCmd, callCmd CallCmd) (*cobra.Command, error) {
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
	)

	return cmd, nil
}

func NewCmdRoot(ctx context.Context, cfg Config) (*cobra.Command, error) {
	wire.Build(
		ProviderCmdRoot,
		ProviderLsCmd,
		ProviderEmptyJSONCmd,
		ProviderToJSONCmd,
		ProviderToBinaryCmd,
		ProviderCallCmd,
		base,
	)

	return nil, nil
}
