package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thorstenpfister/gonyt/internal/nytapi"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var FlagVerbose bool
var FlagJSONOutput bool
var FlagApiKey string

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

	rootCmd.PersistentFlags().BoolVarP(&FlagVerbose, "verbose", "v", false, "Output verbose infos.")
	rootCmd.PersistentFlags().BoolVarP(&FlagJSONOutput, "json", "j", false, "Output in plain JSON instead of formatted overview.")
	rootCmd.PersistentFlags().StringVarP(&FlagApiKey, "apikey", "a", "", "Your key for the New York Times API.")
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
		if FlagVerbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Returns the relevant API Key in order of preference:
// CLI supplied > config file
func preferredApiKey() (*string, error) {
	if FlagApiKey != "" {
		return &FlagApiKey, nil
	}

	configAPIKey := viper.GetString("APIKEY")
	if configAPIKey != "" {
		return &configAPIKey, nil
	}

	return nil, fmt.Errorf("no API key provided via CLI or config file")
}

func printJSON(articles *[]nytapi.Article) error {
	json, err := json.Marshal(articles)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON")
	}

	fmt.Println(string(json))
	return nil
}
