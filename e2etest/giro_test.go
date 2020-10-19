package e2etest_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	cmdtest "github.com/google/go-cmdtest"
	"github.com/pkg/errors"
	"github.com/rerost/giro/e2etest/dummyserver"
	"github.com/rerost/giro/runner/giro"
)

const (
	TestPort = 5000
)

func GiroCmd() int {
	port, closer, err := startServer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer closer()

	ctx := context.Background()
	cmd, err := giro.NewCmdRoot(ctx, giro.Config{
		ReflectionServer: fmt.Sprintf("localhost:%s", port),
	}, "", "")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func startServer() (string, func(), error) {
	port := TestPort
	closer, err := dummyserver.Run(fmt.Sprint(port))
	if err != nil {
		return "", func() {}, errors.WithStack(err)
	}

	return fmt.Sprintf("%v", port), closer, nil
}

func TestGiro(t *testing.T) {
	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["giro"] = cmdtest.InProcessProgram("giro", GiroCmd)
	ts.Run(t, false)
}
