package handler

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"githubtask/internal/github/model"
	"log"
	"testing"
)

func TestGetRepos(t *testing.T) {
	var githubRepo GithubRepo
	repoData, err := githubRepo.GetRepos(0, "", "", NewRequestApiClientMock())
	assert.NoError(t, err)
	assert.NotEmpty(t, repoData)
	repoData, err = githubRepo.GetRepos(0, "", "", NewRequestApiClientError())
	assert.Error(t, err)
	repoData, err = githubRepo.GetRepos(0, "", "", NewRequestApiClientEmptyData())
	assert.Error(t, err)
	assert.Empty(t, repoData)
}

type RequestApiClientMock struct {
}

func NewRequestApiClientMock() *RequestApiClientMock {
	return &RequestApiClientMock{}
}

func (RequestApiClientMock) RequestGetApi(url string) ([]byte, error) {
	var responseMock model.GitHubRepoResponse
	err := gofakeit.Struct(&responseMock)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	responseData, _ := json.Marshal(responseMock)
	return responseData, nil
}

type RequestApiClientError struct {
}

func NewRequestApiClientError() *RequestApiClientError {
	return &RequestApiClientError{}
}

func (RequestApiClientError) RequestGetApi(url string) ([]byte, error) {
	var responseMock model.GitHubRepoResponse
	err := gofakeit.Struct(&responseMock)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	responseData, _ := json.Marshal(responseMock)
	return responseData, assert.AnError
}

type RequestApiClientEmptyData struct {
}

func NewRequestApiClientEmptyData() *RequestApiClientEmptyData {
	return &RequestApiClientEmptyData{}
}

func (RequestApiClientEmptyData) RequestGetApi(url string) ([]byte, error) {
	return nil, nil
}
