package main

import (
	"fmt"
	"os"

	"github.com/rerost/giro/e2etest/dummyserver"
)

func main() {
	closer, err := dummyserver.Run("5000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run server: %v", err)
		os.Exit(1)
	}
	defer closer()

	// ここでサーバを起動しっぱなしにする
	select {}
}
