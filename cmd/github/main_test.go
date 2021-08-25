package main

import (
	"github.com/stretchr/testify/assert"
	"githubtask/internal/github/api/handler"
	"githubtask/pkg/http_helper"
	"testing"
)

func Test_Main(t *testing.T) {
	isConnected:=http_helper.InternetConnected()
	assert.True(t, isConnected,"internet is not connected")
	var githubRepo handler.GithubRepo
	requestApi := http_helper.NewRequestApiClient()
	data, err := githubRepo.GetRepos(5, "golang", "2000-01-01", requestApi)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}
