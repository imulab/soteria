package main

import (
	"fmt"
	"github.com/imulab/soteria/app/authorize"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use: "soteria",
	Short: "OAuth 2.0 / Open ID Connect 1.0 Platform",
}

func main() {
	rootCommand.AddCommand(authorize.ApiCommand())
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
