package main

import (
	"fmt"
	"os"

	"github.com/rerost/giro/runner/giro"
)

func main() {
	if err := giro.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
