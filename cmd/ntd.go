package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ntd",
	Short: "NTD automates your daily todo list",
	Long: `NTD will automatically create todo list items in your daily journal 
			based on the unresolved todo items from the previous journal entry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ntd run, args: %v\n", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
