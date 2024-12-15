/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kzankpe/go-projects/tmdb-cli/helper"
	"github.com/spf13/cobra"
)

var movieType string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmdb-cli",
	Short: "CLI to fetch data from The Movie Database (TMSB) and display it in the terminal.",
	Long: `The application run from the command line, and is able to pull and show the popular, top-rated, upcoming and now playing movies from the TMDB API.\n
The user can specify the type of movies they want to see by passing a command line argument to the CLI tool.`,
	// Args: cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if movieType == "" {
			fmt.Println("Please provide a category using --type (e.g., popular, top_rated, upcoming, now_playing).")
			return
		}
		fmt.Println(movieType)
		endpoint, err := helper.ValidateTypeInput(movieType)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		movies, err := helper.FetchMovies(endpoint)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Movies in category '%s':\n", movieType)
		for _, movie := range movies {
			fmt.Println("-", movie.Title)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tmdb-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&movieType, "type", "t", "", "Type of the movie")
}
