package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thorstenpfister/gonyt/nytapi"
)

var bookreviewsFlagCategory string
var bookreviewsFlagSearchTerm string

var bookreviewsCmd = &cobra.Command{
	Use:   "bookreviews",
	Short: "Fetch book reviews from the New York Times for a search term and category.",
	Long: `Fetch book reviews from the New York Times for a search term and category.

	Sections include: 	
		author, isbn, title

	Example usage:
		gonyt bookreviews -c author -t "Michelle Obama"
		gonyt bookreviews -c title -t "Finders Keepers"`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newCLIClient()
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}
		ctx := context.Background()
		category := nytapi.BookReviewsCategory(bookreviewsFlagCategory)

		bookReviews, err := client.FetchBookReviews(ctx, category, bookreviewsFlagSearchTerm)
		if err != nil {
			fmt.Println("Error calling New York Times API!", err)
			return
		}

		printBookReviews(bookReviews)
	},
}

func init() {
	rootCmd.AddCommand(bookreviewsCmd)

	bookreviewsCmd.Flags().StringVarP(&bookreviewsFlagCategory, "category", "c", "", "Book review category to search for.")
	bookreviewsCmd.Flags().StringVarP(&bookreviewsFlagSearchTerm, "term", "t", "", "Book review term to search for.")
	bookreviewsCmd.MarkFlagRequired("category")
	bookreviewsCmd.MarkFlagRequired("searchTerm")
}
