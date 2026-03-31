package cmd

import (
	"fmt"
	"os"

	"github.com/bnassif/jquery-proxy/pkg"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Version: pkg.Version,
	Use:     "jquery-proxy",
	Short:   "",
	Long:    "",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(runCmd)
}
