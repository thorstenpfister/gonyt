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
func print(articles *[]nytapi.Article, updateTime *time.Time) {
	if flagJSONOutput {
		err := printJSON(articles)
		if err != nil {
			fmt.Println("Error printing JSON!", err)
			return
		}
	} else {
		printArticles(articles, updateTime)
	}
}

// Handles printing of articles as JSON array
func printJSON(articles *[]nytapi.Article) error {
	json, err := json.Marshal(articles)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON")
	}

	fmt.Println(string(json))
	return nil
}

// Handles opinionated printing of articles
func printArticles(articles *[]nytapi.Article, updateTime *time.Time) {
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
