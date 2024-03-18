package commandHandlers

import (
	"os"
	"slices"
	"strings"

	"github.com/Majunga/go_cli/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/yargevad/filepathx"
)

var pathsToIgnore = []string{
	"node_modules",
	".git",
	".vs",
}

func CreateCleanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Remove Bin and Obj folders recursively in a directory (ignores node_modules)",
		Run:   cleanCommandHandler,
	}
}

func cleanCommandHandler(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	// doing glob on both bin and obj folders and then removing the ignored paths is very inefficient
	// needs to be improved in the future maybe, who knows
	cmd.Printf("Cleaning directory %s", wd)
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

	reducedDirectories := utils.Reduce(directories, func(s string) bool {
		for _, ignore := range pathsToIgnore {
			if strings.Contains(s, ignore) {
				return true
			}
		}

		return false
	})

	if len(reducedDirectories) == 0 {
		cmd.Println("Nothing to delete :D")
		return
	}

	cmd.Println("Found the following directories to delete:")
	for _, d := range reducedDirectories {
		cmd.Println(d)
	}

	cmd.Println("Are you sure you want to delete these directories? [y/N]")
	var keyByte []byte = []byte{}

	cmd.InOrStdin().Read(keyByte)

	key, _ := utils.Read()

	if strings.ToLower(key) != "y" {
		return
	}

	for _, d := range reducedDirectories {
		if err := os.RemoveAll(d); err != nil {
			cmd.PrintErr(err)
			return
		}
	}

	cmd.Println("Successfully deleted all bin and obj folders!")
}
