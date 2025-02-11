/*
Copyright © 2025 Aditya Chaphekar
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aditya-chaphekar/create-util/models"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project with the selected framework and default configuration.",
	Long: `The init command sets up a new project by selecting a framework and applying a default configuration. 
It scaffolds the necessary project structure and files, allowing for a quick start. 

Users can choose from various frameworks (e.g., React, Vue, Svelte, Angular) through an interactive selection 
or by passing options via the command line. Additional configurations, such as package managers and 
project settings, can also be included.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = models.GetProjectName()
		selectedType, err := models.GetProjectType()
		if err != nil {
			fmt.Println("Failed to select project type:", err)
			os.Exit(1)
		}
		fmt.Print("Selected project type:", selectedType)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
