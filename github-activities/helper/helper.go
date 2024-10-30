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
	Action  string   `json:action`
	Ref     string   `json:"ref"`
	Commits []Commit `json:commits`
}

type Commit struct {
	Sha     string `json:sha`
	Message string `json:message`
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
	for _, event := range events {
		switch event.Type {
		case "IssuesEvent":
			fmt.Printf("- %s a new issue in %s\n", event.Payload.Action, event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("- %s a Pull Request in %s\n", event.Payload.Action, event.Repo.Name)
		case "PushEvent":
			fmt.Printf("- Pushed %d commit to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("- Created %s in %s\n", event.Payload.Ref, event.Repo.Name)
		}
	}
}
