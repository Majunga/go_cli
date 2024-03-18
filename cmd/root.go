package cmd

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yargevad/filepathx"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_cli",
	Short: "CLI tool to help me with automating few things",
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove Bin and Obj folders recursively in a directory (ignores node_modules)",
	Run:   cleanCommandHandler,
}

func cleanCommandHandler(cmd *cobra.Command, args []string) {
	binDirectories, err := filepathx.Glob("./**/bin")
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	objDirectories, err := filepathx.Glob("./**/obj")

	if err != nil {
		cmd.PrintErr(err)
		return
	}

	directories := append(binDirectories, objDirectories...)
	slices.Sort(directories) // makes it easier on the user to check before deleting

	if len(directories) == 0 {
		cmd.Println("Nothing to delete :D")
		return
	}

	cmd.Println("Found the following directories to delete:")
	for _, d := range directories {
		cmd.Println(d)
	}

	cmd.Println("Are you sure you want to delete these directories? [y/N]")
	var keyByte []byte = []byte{}

	cmd.InOrStdin().Read(keyByte)

	key, _ := read()

	if strings.ToLower(key) != "y" {
		return
	}

	for _, d := range directories {
		if err := os.RemoveAll(d); err != nil {
			cmd.PrintErr(err)
			return
		}
	}

	cmd.Println("Successfully deleted all bin and obj folders!")
}

func read() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')

		if len(strings.TrimSpace(text)) > 0 {
			return strings.TrimSpace(text), err
		}
	}
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
	rootCmd.AddCommand(cleanCmd)
}
