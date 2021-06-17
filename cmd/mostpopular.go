package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thorstenpfister/gonyt/nytapi"
)

var mostPopularFlagCategory string
var mostPopularFlagPeriod int

// mostpopularCmd represents the mostpopular command
var mostpopularCmd = &cobra.Command{
	Use:   "mostpopular",
	Short: "Fetch the most popular stories from the New York Times for a given category and time period.",
	Long: `Fetch the most popular stories from the New York Times for a given category and time period.
	
	Sections include:
		emailed, shared, viewed

	Time periods (in days) include:
		1, 7, 30

	Example usage:
		gonyt mostpopular -c emailed -p 7
		gonyt mostpopular -c viewed -p 30
	`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newCLIClient()
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}
		ctx := context.Background()
		category := nytapi.MostPopularCategory(mostPopularFlagCategory)
		period := nytapi.MostPopularPeriod(mostPopularFlagPeriod)

		articles, err := client.FetchMostPopularArticles(ctx, category, period)
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}

		printPopularArticles(articles)
	},
}

func init() {
	rootCmd.AddCommand(mostpopularCmd)

	mostpopularCmd.Flags().StringVarP(&mostPopularFlagCategory, "category", "c", "", "Most popular articles category to be fetched.")
	mostpopularCmd.MarkFlagRequired("category")
	mostpopularCmd.Flags().IntVarP(&mostPopularFlagPeriod, "period", "p", 1, "Most popular articles time period to be fetched.")
	mostpopularCmd.MarkFlagRequired("period")
}
