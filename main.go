package main

import (
	"fmt"
	"github.com/imulab/soteria/app/authorize"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use: "soteria",
	Short: "OAuth 2.0 / Open ID Connect 1.0 Platform",
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	rootCommand.AddCommand(authorize.ApiCommand())
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
