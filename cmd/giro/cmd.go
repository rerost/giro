package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	ReflectionServer string `mapstructure:"reflection-server"`
	RpcServer        string `mapstructure:"rpc-server"`
	Debug            bool
	Verbose          bool
}

func Run() error {
	ctx := context.Background()
	cfg, err := NewConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	l, err := NewLogger(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	defer l.Sync()

	zap.ReplaceGlobals(l)
	cfgJSON, _ := json.Marshal(cfg)
	zap.L().Debug("config", zap.String("config", string(cfgJSON)))

	cmd, err := NewCmdRoot(ctx, cfg)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := cmd.Execute(); err != nil {
		if cfg.Debug {
			fmt.Printf("%+v\n", err)
		}
		return errors.WithStack(err)
		// zap.L().Debug("error", zap.String("stack trace", fmt.Sprintf("%+v\n", err)))
	}
	return nil
}

func NewLogger(cfg Config) (*zap.Logger, error) {
	zcfg := zap.NewProductionConfig()
	if cfg.Debug {
		zcfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	if cfg.Verbose {
		zcfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	l, err := zcfg.Build()
	return l, errors.WithStack(err)
}

func NewConfig() (Config, error) {
	pflag.StringP("reflection-server", "", "localhost:5000", "")
	pflag.StringP("rpc-server", "", "", "")
	pflag.BoolP("verbose", "", false, "")
	pflag.BoolP("debug", "", false, "")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	viper.BindPFlags(pflag.CommandLine)

	var cfg Config
	pflag.Parse()
	err := viper.Unmarshal(&cfg)
	return cfg, errors.WithStack(err)
}
