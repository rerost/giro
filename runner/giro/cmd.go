package giro

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	ReflectionServer string `mapstructure:"reflection-server"`
	RpcServer        string `mapstructure:"rpc-server"`
	Debug            bool
	Verbose          bool
	Metadata         string
}

func Run(version Version, revision Revision) error {
	ctx := context.Background()
	cmd, err := NewCmdRoot(ctx, version, revision)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := cmd.Execute(); err != nil {
		zap.L().Debug("error", zap.String("stack trace", fmt.Sprintf("%+v\n", err)))
		return errors.WithStack(err)
	}
	return nil
}
