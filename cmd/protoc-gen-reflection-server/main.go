package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rerost/giro/runner/genreflectionserver"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		Error(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		Error(err)
	}

	resp, err := genreflectionserver.Run(req)
	if err != nil {
		Error(err)
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		Error(err)
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		Error(err)
	}
}

func Debug(vs ...interface{}) {
	fmt.Fprintln(os.Stderr, vs...)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
