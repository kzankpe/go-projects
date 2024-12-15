/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kzankpe/go-projects/tmdb-cli/cmd"
	"github.com/kzankpe/go-projects/tmdb-cli/helper"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get the API key from the environment variable
	helper.ApiKey = os.Getenv("TMDB_API_KEY")
	if helper.ApiKey == "" {
		fmt.Println("TMDB_API_KEY is not set in the environment")
		return
	}
	cmd.Execute()
}
