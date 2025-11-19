package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "x-ui",
	Short: "A powerful Xray web panel",
	Long:  `x-ui is a web panel for Xray with a focus on user experience.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
