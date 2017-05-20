package main

import (
	"fmt"
	"os"

	"github.com/coreos/dex/cmd/dex"
)

func main() {
	if err := dex.GetCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
