package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
		section := nytapi.TopStoriesSection(Section)

		apiKey, err := preferredApiKey()
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}
		if FlagVerbose {
			fmt.Println("Using api key:", apiKey)
		}

		httpClient := http.Client{
			Timeout: 15 * time.Second,
		}

		client := nytapi.NewClient(&httpClient, *apiKey)
		ctx := context.Background()

		articles, updateTime, err := client.FetchTopStories(ctx, section)
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}

		if FlagJSONOutput {
			err = printJSON(articles)
			if err != nil {
				fmt.Println("Error printing JSON!", err)
				return
			}
		} else {
			printTopstories(articles, updateTime)
		}
	},
}

func init() {
	rootCmd.AddCommand(topstoriesCmd)

	topstoriesCmd.Flags().StringVarP(&Section, "section", "s", "", "Top stories section to be fetched.")
	topstoriesCmd.MarkFlagRequired("section")
}

func printTopstories(articles *[]nytapi.Article, updateTime *time.Time) {
	fmt.Println("Last update:", updateTime)
	fmt.Println()

	for _, article := range *articles {
		fmt.Println(article.Title)
		if article.Abstract != "" {
			fmt.Println("\t", article.Abstract)
		}
		fmt.Println("\t", article.Url)
	}
}
