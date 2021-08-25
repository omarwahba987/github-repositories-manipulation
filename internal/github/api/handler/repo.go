package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"githubtask/internal/github/model"
	httpHelper "githubtask/pkg/http_helper"
	"log"
	"strconv"
	"time"
)

type Github interface {
	GetRepos(limit int, key string, dateStr string) (model.GitHubRepoResponse, error)
}
type GithubRepo struct {
}

func (g *GithubRepo) GetRepos(limit int, key string, dateStr string, requestApi httpHelper.RequestApi) (model.GitHubRepoResponse, error) {
	date, _ := time.Parse(model.DateLayoutFormat, dateStr)
	limitStr := strconv.Itoa(limit)
	restUrl := fmt.Sprintf("search/repositories?q=%s+language:assembly&sort=stars&order=desc&per_page=%v", key, limitStr)
	url := model.GithubBaseLink + restUrl
	var responseObject model.GitHubRepoResponse
	responseData, err := requestApi.RequestGetApi(url)
	if err != nil {
		log.Println(err)
		return model.GitHubRepoResponse{}, err
	}
	if responseData == nil {
		log.Println("empty data received")
		return model.GitHubRepoResponse{}, errors.New("empty data received")
	}
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Println(err)
		return model.GitHubRepoResponse{}, err
	}
	for i := 0; i < len(responseObject.Items); i++ {
		if date.After(responseObject.Items[i].CreatedAt) {
			responseObject.Items = append(responseObject.Items[:i], responseObject.Items[i+1:]...)
		}
	}
	return responseObject, nil
}
