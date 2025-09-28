package app

import (
	"fmt"
	"os"
)

var Version string = "dev"

func PrintVersion() {
	fmt.Printf("jquery-proxy version %s\n", Version)
	os.Exit(0)
}
