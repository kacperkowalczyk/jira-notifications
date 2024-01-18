package main

import (
	"context"
	"fmt"
	"time"

	"encoding/json"
	"io"
	"log"
	"os"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/gen2brain/beeep"
)

func NewIssueNotification(ticketCount int) {
	message := fmt.Sprintf("Number of unassigned issues: %d", ticketCount)
	beeep.Notify("New unassigned issues found", message, "")
	fmt.Println("New issues found")
}

type Config struct {
	JiraURL  string `json:"jira_url"`
	Token    string `json:"token"`
	JQL      string `json:"jql"`
	Interval int    `json:"interval"`
}

func GetConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func CheckForNewIssues(jiraClient *jira.Client, config Config) {

	issues, _, err := jiraClient.Issue.Search(context.Background(), config.JQL, nil)
	if err != nil {
		panic(err)
	}

	if len(issues) > 0 {
		NewIssueNotification(len(issues))
	}
}

func main() {
	config := GetConfig()
	tp := jira.BearerAuthTransport{
		Token: config.Token,
	}
	jiraClient, err := jira.NewClient(config.JiraURL, tp.Client())
	if err != nil {
		panic(err)
	}

	u, _, err := jiraClient.User.GetSelf(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Logged in as: %v\n", u.EmailAddress)

	for {
		CheckForNewIssues(jiraClient, config)

		time.Sleep(time.Duration(config.Interval) * time.Minute)
	}

}
