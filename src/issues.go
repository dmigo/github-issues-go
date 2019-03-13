package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const BASE_URI = "https://api.github.com"
const OWNER = "gotify"
const PROJECT_NAME = "server"
const BUG_LABEL = "a:bug"

func GetUrl() string {
	var Url *url.URL
	Url, err := url.Parse(BASE_URI)

	if err != nil {
		panic(fmt.Sprintf("bad url: %s", BASE_URI))
	}

	Url.Path += "/repos/" + OWNER + "/" + PROJECT_NAME + "/issues"
	parameters := url.Values{}
	parameters.Add("labels", BUG_LABEL)
	Url.RawQuery = parameters.Encode()

	return Url.String()
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func listIssues() ([]*Issue, error) {
	var url = GetUrl()
	var issues []*Issue
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	decodeErr := json.NewDecoder(r.Body).Decode(&issues)

	return issues, decodeErr
}

func main() {
	issues, err := listIssues()

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("Issues")
		fmt.Println("------")
		for _, issue := range issues {
			fmt.Printf("[%v] %v\n", *issue.ID, *issue.Title)
			fmt.Printf("by %v\n", *issue.User.Login)
			fmt.Println("------")
		}
	}
}
