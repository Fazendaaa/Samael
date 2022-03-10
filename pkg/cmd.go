package samael

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func checkPath(path string) error {
	dir, fail := os.Stat(path)

	if fail != nil {
		return fmt.Errorf(`invalid path: %s
	given error: %q`, path, fail)
	}
	if !dir.IsDir() {
		return fmt.Errorf("%q is not a directory", dir.Name())
	}

	return nil
}

func ValidPath(command string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, params []string) error {
		if 1 != len(params) {
			return fmt.Errorf(`missing project path!
If you want to %s a project in the current path, you can run it again as:
	shojo %s .
Otherwise you may try to:
	shojo %s /absolute/or/relative/project/path`, command, command, command)
		}

		return checkPath(params[0])
	}
}

func ValidateProjectPath(projectPath string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, params []string) error {
		if 1 > len(params) {
			return fmt.Errorf("requires at least 1 arg(s), only received 0")
		}

		if "" == projectPath {
			path, fail := os.Getwd()

			if nil != fail {
				fmt.Printf("%v;\ncould not read current directory", fail)
			}

			projectPath = path
		}

		return checkPath(projectPath)
	}
}

type toConsume func(params []string, projectPath string) (channel chan ResponseChannel, fail error)

func RunCmd(projectPath *string, message string, function toConsume) func(cmd *cobra.Command, params []string) {
	return func(cmd *cobra.Command, params []string) {
		spinner, fail := CreateSpinner(message, "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel, fail := function(params, *projectPath)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		fail = ConsumeChannel(params, spinner, resultChannel)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			KillSpinner(spinner, false)

			return
		}

		KillSpinner(spinner, true)
	}
}
