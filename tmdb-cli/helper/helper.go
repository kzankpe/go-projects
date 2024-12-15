package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const BaseUrl = "https://api.themoviedb.org/3/movie/"

var ApiKey string

type Movie struct {
	Title string `json:"title"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

func ValidateTypeInput(input string) (string, error) {
	validInput := map[string]string{
		"playing":  "now_playing",
		"popular":  "popular",
		"top":      "top_rated",
		"upcoming": "upcoming",
	}

	if endpoint, exist := validInput[input]; exist {
		return endpoint, nil
	}

	return "", errors.New("invalid filter value. Allowed values are: now_playing, top_rated, popular, upcoming")
}

func FetchMovies(endpoint string) ([]Movie, error) {
	url := fmt.Sprintf("%s%s", BaseUrl, endpoint)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+ApiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", response.StatusCode, string(body))
	}

	var movieResponse MovieResponse
	if err := json.NewDecoder(response.Body).Decode(&movieResponse); err != nil {
		return nil, err
	}

	return movieResponse.Results, nil
}
