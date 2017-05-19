package main

import (
	"fmt"
	"os"

	"github.com/coreos/dex/cmd/dex"
)

func main() {
	fmt.Println("HEY ROBIN")
	if err := dex.GetCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
