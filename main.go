package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	defer panicRecover()

	args := os.Args

	githubAccount := fmt.Sprintf("https://api.github.com/users/%s/events?per_page=100", args[1])

	res, err := http.Get(githubAccount)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNotFound:

	}

	var githubResponse []githubResponse
	err = json.NewDecoder(res.Body).Decode(&githubResponse)
	if err != nil {
		panic(err)
	}

	for _, v := range githubResponse {
		printEvent(event(v.Type), v)
	}
}

func panicRecover() {
	if r := recover(); r != nil {
		fmt.Printf("Error recovered: %v\n", r)
	}
}

type githubResponse struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

type event string

const (
	CommitCommentEvent            event = "CommitCommentEvent"
	CreateEvent                   event = "CreateEvent"
	DeleteEvent                   event = "DeleteEvent"
	ForkEvent                     event = "ForkEvent"
	GollumEvent                   event = "GollumEvent"
	IssueCommentEvent             event = "IssueCommentEvent"
	IssuesEvent                   event = "IssuesEvent"
	MemberEvent                   event = "MemberEvent"
	PublicEvent                   event = "PublicEvent"
	PullRequestEvent              event = "PullRequestEvent"
	PullRequestReviewEvent        event = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent event = "PullRequestReviewCommentEvent"
	PullRequestReviewThreadEvent  event = "PullRequestReviewThreadEvent"
	PushEvent                     event = "PushEvent"
	ReleaseEvent                  event = "ReleaseEvent"
	SponsorshipEvent              event = "SponsorshipEvent"
	WatchEvent                    event = "WatchEvent"
)

func printEvent(event event, data githubResponse) {
	switch event {
	case CommitCommentEvent:
		fmt.Printf("Commented on a commit in %s\n", data.Repo.Name)
	case CreateEvent:
		fmt.Printf("Created %s in %s\n", data.Payload.RefType, data.Repo.Name)
	case DeleteEvent:
		fmt.Printf("Deleted %s in %s\n", data.Payload.RefType, data.Repo.Name)
	case ForkEvent:
		fmt.Printf("Forked %s\n", data.Repo.Name)
	case GollumEvent:
		fmt.Printf("Updated the wiki in %s\n", data.Repo.Name)
	case IssueCommentEvent:
		fmt.Printf("Commented on issue in %s\n", data.Repo.Name)
	case IssuesEvent:
		fmt.Printf("Opened a new issue in %s\n", data.Repo.Name)
	case MemberEvent:
		fmt.Printf("Added a member to %s\n", data.Repo.Name)
	case PublicEvent:
		fmt.Printf("Made %s public\n", data.Repo.Name)
	case PullRequestEvent:
		fmt.Printf("created pull request on %s\n", data.Repo.Name)
	case PullRequestReviewEvent:
		fmt.Printf("Reviewed pull request on %s\n", data.Repo.Name)
	case PullRequestReviewCommentEvent:
		fmt.Printf("Comented on pull request on %s\n", data.Repo.Name)
	case PullRequestReviewThreadEvent:
		fmt.Printf("Started a review thread in %s\n", data.Repo.Name)
	case PushEvent:
		fmt.Printf("Pushed %d commits to %s\n", len(data.Payload.Commits), data.Repo.Name)
	case ReleaseEvent:
		fmt.Printf("Released a new version in %s\n", data.Repo.Name)
	case SponsorshipEvent:
		fmt.Printf("Sponsored a project in %s\n", data.Repo.Name)
	case WatchEvent:
		fmt.Printf("Starred %s\n", data.Repo.Name)
	}
}
