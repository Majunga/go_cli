package commandHandlers

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yargevad/filepathx"
)

type CommandHandler struct {
	*cobra.Command
}

func CreateGlobExeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "globExe pattern command",
		Short: "Glob for files and excute command on each file it finds in a directory\n\t\tExample: globExe \"**/*.zig\" \"zig test %s\"",
		Args:  cobra.ExactArgs(2),
		Run:   globExeCommandHandler,
	}
}

func globExeCommandHandler(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	pattern := args[0]
	command := args[1]

	cmdHandler := CommandHandler{cmd}
	fmt.Printf("Globbing for files in directory%s\n", wd)
	files, err := filepathx.Glob(pattern)
	if err != nil {
		cmdHandler.PrintErr(err)
		return
	}

	fmt.Printf("Found %d files\n", len(files))
	for _, file := range files {
		cmdHandler.executeCommand(fmt.Sprintf(command, file))
	}
}

func (commandHandler *CommandHandler) executeCommand(fullCommand string) {
	fmt.Printf("Executing command: %s\n", fullCommand)
	commands := strings.Split(fullCommand, "&&")
	for _, command := range commands {
		commandParts, err := trim(command)
		if err != nil {
			commandHandler.PrintErr(err)
			return
		}

		var cmd *exec.Cmd

		if len(commandParts) > 1 {
			args := commandParts[1:]
			cmd = exec.Command(commandParts[0], args...)
		} else {
			cmd = exec.Command(commandParts[0])
		}

		stdout, err := cmd.Output()

		if err != nil {
			stdErr := err.(*exec.ExitError)
			fmt.Printf("%s", stdErr.Stderr)
		}

		fmt.Printf("%s", stdout)
	}
}

// Splitter splits a string command into command and arguments
func splitter(s string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ' ' // space
	fields, err := r.Read()

	return fields, err
}

func trim(commandString string) ([]string, error) {
	result := []string{}
	splitCommand, err := splitter(strings.TrimSpace(commandString))
	for _, s := range splitCommand {
		result = append(result, strings.TrimSpace(s))
	}

	return result, err
}
