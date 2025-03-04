package cmd

import (
	"fmt"

	"github.com/0biglife/dev-quickstart/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-quickstart",
	Short: "🚀 A CLI for setting up dev environments",
	Long:  "This CLI helps you quickly scaffold frontend, backend, and CI/CD setups.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintInfo("📢 Use 'dev-quickstart --help' for more information about a command.")
	},
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Select which project to set up (frontend, backend, ci/cd)",
	Run: func(cmd *cobra.Command, args []string) {
		var choice string
		utils.PrintPrompt("Please enter which project to set up (frontend/backend/ci-cd): ")
		fmt.Scanln(&choice)

		switch choice {
		case "frontend":
			setupFrontend()
		case "backend":
			setupBackend()
		case "ci-cd":
			setupCI_CD()
		default:
			utils.PrintError("Invalid choice! Please enter 'frontend', 'backend', or 'ci-cd'.")
		}
	},
}

func Execute() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.Execute()
}
