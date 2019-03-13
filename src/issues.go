package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func GetUrl(configuration Configuration, arguments Arguments) string {
	var Url *url.URL
	Url, err := url.Parse(*configuration.BaseUrl)

	if err != nil {
		panic(fmt.Sprintf("bad url: %s", configuration.BaseUrl))
	}

	Url.Path += "/repos/" + *arguments.Owner + "/" + *arguments.Project + "/issues"
	parameters := url.Values{}
	parameters.Add("labels", *arguments.Labels)
	Url.RawQuery = parameters.Encode()

	return Url.String()
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func listIssues(configuration Configuration, arguments Arguments) ([]*Issue, error) {
	var url = GetUrl(configuration, arguments)
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

	arguments := ReadArguments()
	configuration := GetConfiguration()

	issues, err := listIssues(configuration, arguments)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("Issues:")
		fmt.Println()
		for index, issue := range issues {
			fmt.Printf("%v.\t[%v] %v\n", index, *issue.ID, *issue.Title)
			fmt.Printf("\tby %v\n", *issue.User.Login)
			fmt.Println()
		}
	}
}
