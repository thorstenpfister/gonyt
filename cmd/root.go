package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/thorstenpfister/gonyt/nytapi"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var flagVerbose bool
var flagJSONOutput bool
var flagApiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonyt",
	Short: "A command-line interface for the New York Times API.",
	Long:  `A command-line interface for the New York Times API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Output verbose infos.")
	rootCmd.PersistentFlags().BoolVarP(&flagJSONOutput, "json", "j", false, "Output in plain JSON instead of formatted overview.")
	rootCmd.PersistentFlags().StringVarP(&flagApiKey, "apikey", "a", "", "Your key for the New York Times API.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".gonyt" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".gonyt")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if flagVerbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Returns a client suitable for any CLI command
func newCLIClient() (*nytapi.Client, error) {
	apiKey, err := preferredApiKey()
	if err != nil {
		return nil, err
	}
	if flagVerbose {
		fmt.Println("Using api key:", apiKey)
	}

	httpClient := http.Client{
		Timeout: 15 * time.Second,
	}

	client := nytapi.NewClient(&httpClient, *apiKey)
	return &client, nil
}

// Returns the relevant API Key in order of preference:
// CLI supplied > config file
func preferredApiKey() (*string, error) {
	if flagApiKey != "" {
		return &flagApiKey, nil
	}

	configAPIKey := viper.GetString("APIKEY")
	if configAPIKey != "" {
		return &configAPIKey, nil
	}

	return nil, fmt.Errorf("no API key provided via CLI or config file")
}

// Handles general printing based on CLI flags
func printArticles(articles *[]nytapi.Article, updateTime *time.Time) {
	if flagJSONOutput {
		err := printJSONArticles(articles)
		if err != nil {
			fmt.Println("Error printing JSON!", err)
			return
		}
	} else {
		printArticlesCLI(articles, updateTime)
	}
}

// Handles general printing of popular articles based on CLI flags
func printPopularArticles(articles *[]nytapi.PopularArticle) {
	if flagJSONOutput {
		err := printJSONPopularArticles(articles)
		if err != nil {
			fmt.Println("Error printing JSON!", err)
			return
		}
	} else {
		printPopularArticlesCLI(articles)
	}
}

// Handles general printing based on CLI flags
func printBookReviews(bookReviews *[]nytapi.BookReview) {
	if flagJSONOutput {
		err := printJSONBookReviews(bookReviews)
		if err != nil {
			fmt.Println("Error printing JSON!", err)
			return
		}
	} else {
		printBookReviewsCLI(bookReviews)
	}
}

// Handles printing of articles as JSON array
func printJSONArticles(articles *[]nytapi.Article) error {
	json, err := json.Marshal(articles)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON")
	}

	fmt.Println(string(json))
	return nil
}

// Handles printing of popular articles as JSON array
func printJSONPopularArticles(articles *[]nytapi.PopularArticle) error {
	json, err := json.Marshal(articles)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON")
	}

	fmt.Println(string(json))
	return nil
}

// Handles printing of book reviews as JSON array
func printJSONBookReviews(bookReviews *[]nytapi.BookReview) error {
	json, err := json.Marshal(bookReviews)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON")
	}

	fmt.Println(string(json))
	return nil
}

// Handles opinionated printing of articles
func printArticlesCLI(articles *[]nytapi.Article, updateTime *time.Time) {
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

// Handles opinionated printing of popular articles
func printPopularArticlesCLI(articles *[]nytapi.PopularArticle) {
	for _, article := range *articles {
		fmt.Println(article.Title)
		if article.Abstract != "" {
			fmt.Println("\t", article.Abstract)
		}
		fmt.Println("\t", article.URL)
	}
}

// Handles opinionated printing of book reviews
func printBookReviewsCLI(bookReviews *[]nytapi.BookReview) {
	for _, bookReview := range *bookReviews {
		fmt.Println(bookReview.BookAuthor, " - ", bookReview.BookTitle)
		if bookReview.Summary != "" {
			fmt.Println("\t", bookReview.Summary)
		}
		fmt.Println("\t", bookReview.URL)
	}
}
