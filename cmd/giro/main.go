package main

import (
	"fmt"
	"os"

	"github.com/rerost/giro/runner/giro"
)

var (
	Version  = "Empty"
	Revision = "Empty"
)

func main() {
	if err := giro.Run(giro.Version(Version), giro.Revision(Revision)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
