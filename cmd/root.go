package cmd

import (
	"os"

	"github.com/Majunga/go_cli/cmd/commandHandlers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_cli",
	Short: "CLI tool to help me with automating few things",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commandHandlers.CreateCleanCommand())
	rootCmd.AddCommand(commandHandlers.CreateGlobExeCommand())
}
