package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ikovic/notion-todo-daily/internal/notion"
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

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Page search",
	Long:  "Search the pages from your Notion workspace",
	Run: func(cmd *cobra.Command, args []string) {
		lastPage := notion.GetLastPage(cmd.Root().Context())
		fmt.Printf("Last created page ID: %s\n", lastPage.Id)
	},
}

var getBlocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "List blocks",
	Long:  "List blocks from the given page",
	Run: func(cmd *cobra.Command, args []string) {
		blocks := notion.GetPageBlocks(cmd.Root().Context())
		fmt.Printf("Blocks in the page: %v\n", blocks)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(getBlocksCmd)
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
