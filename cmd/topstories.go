package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thorstenpfister/gonyt/nytapi"
)

var Section string

var topstoriesCmd = &cobra.Command{
	Use:   "topstories",
	Short: "Fetch top stories from a New York Times section.",
	Long: `Fetch top stories from a New York Times section.

	Sections include: 	
		arts, automobiles, books, business, fashion, food, 
		health, home, insider, magazine, movies, nyregion, 
		obituaries, opinion, politics, realestate, science, 
		sports, sundayreview, technology, theater, t-magazine, 
		travel, upshot, us, world

	Example usage:
		gonyt topstories -s opinion		
		gonyt topstories -s magazine`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newCLIClient()
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}
		ctx := context.Background()
		section := nytapi.TopStoriesSection(Section)

		articles, updateTime, err := client.FetchTopStories(ctx, section)
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}

		print(articles, updateTime)
	},
}

func init() {
	rootCmd.AddCommand(topstoriesCmd)

	topstoriesCmd.Flags().StringVarP(&Section, "section", "s", "", "Top stories section to be fetched.")
	topstoriesCmd.MarkFlagRequired("section")
}
