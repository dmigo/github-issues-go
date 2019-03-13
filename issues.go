package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetUrl() string {

	const BASE_URI = "https://api.github.com"
	const OWNER = "gotify"
	const PROJECT_NAME = "server"
	const BUG_LABEL = "a:bug"

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

func main() {
	var url = GetUrl()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		bodyBytes, readErr := ioutil.ReadAll(resp.Body)

		if readErr != nil {
			fmt.Printf("read error: %v\n", err)
		} else {
			bodyString := string(bodyBytes)
			fmt.Printf("result: %v\n", resp.StatusCode)
			fmt.Printf("result: %v\n", bodyString)
		}
	}

	fmt.Println("termingating")
}
