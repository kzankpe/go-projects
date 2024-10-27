package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Activity struct {
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
}

type Repo struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Payload struct {
	Action string `json:action`
	Ref    string `json:"ref"`
}

func GetGithubActivity(username string) ([]Activity, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("User not found. Please check the username spelling.")
	}
	var results []Activity
	err = json.NewDecoder(response.Body).Decode(&results)

	if err != nil {
		return nil, err
	}
	return results, nil
}

func DisplayActivity(events []Activity) {

}