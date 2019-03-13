package main

type Configuration struct {
	BaseUrl *string
}

func GetConfiguration() Configuration {
	url := "https://api.github.com"
	return Configuration{BaseUrl: &url}
}
