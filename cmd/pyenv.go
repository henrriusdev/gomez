/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/hbourgeot/gomez/helpers"

	"github.com/spf13/cobra"
)

// pyenvCmd represents the pyenv command
var pyenvCmd = &cobra.Command{
	Use:   "pyenv",
	Short: "Install Python Version Management",
	Long: `Install pyenv from its website https://pyenv.run.

Provide a version for install, default is 3.9. By default, the path environment variable is configured on .profile file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// variables
		var version, shell, sourceFile string
		// version
		version, _ = cmd.Flags().GetString("version")
		if version == "" {
			version = "3.10"
		}
		fmt.Println("Installing Python version", version)

		forZsh, _ := cmd.Flags().GetBool("zsh")
		forBash, _ := cmd.Flags().GetBool("bash")
		forFish, _ := cmd.Flags().GetBool("fish")

		if forZsh {
			shell = "zsh"
			sourceFile = "~/.zshrc"
		} else if forFish {
			shell = "fish"
			sourceFile = "~/.config/fish/config.fish"
		} else if forBash {
			shell = "bash"
			sourceFile = "~/.bashrc"
		}

		// Call the function to install fnm
		err := helpers.InstallPyenv(shell, sourceFile, version)
		if err != nil {
			fmt.Println("Error installing pyenv")
			fmt.Println(err)
			return
		}

		fmt.Println("Process finished.")
	},
}

func init() {
	rootCmd.AddCommand(pyenvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pyenvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	pyenvCmd.Flags().StringP("version", "v", "", "Install the Node.js version specified")
}
