package main

import (
	"flag"
	"fmt"
	"githubtask/internal/github/api/handler"
	"githubtask/pkg/http_helper"
	"log"
)

func main() {
	if !http_helper.InternetConnected() {
		log.Fatal("no internet connection")
	}
	var (
		keyWordFlag = flag.String("key", "", "key word for filter the data like programing language")
		limitFlag   = flag.Int("limit", 10, "limit for the result data with default 10")
		DateFlag    = flag.String("date", "1990-01-01", "date for filter the received ")
	)

	flag.Parse()
	var githubRepo handler.GithubRepo
	requestApi := http_helper.NewRequestApiClient()
	data, err := githubRepo.GetRepos(*limitFlag, *keyWordFlag, *DateFlag, requestApi)
	if err != nil {
		log.Println(err)
	}

	for _, v := range data.Items {
		fmt.Println(v)
	}
}
