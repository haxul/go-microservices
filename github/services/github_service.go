package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/haxul/go-microservices/github/model"
	"io/ioutil"
	"net/http"
)

type githubService struct{}

var (
	GithubService githubService
)

const (
	token      = "test"
	authHeader = "Authorization"
	apiUrl     = "https://api.github.com"
)

func (g *githubService) CreateRepo(name, description string) (*model.CreateRepoResponse, error) {
	body := model.CreateRepoModel{
		Name:        name,
		Description: description,
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProject:  true,
		HasWiki:     false,
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	request, reqErr := http.NewRequest(http.MethodPost, apiUrl+"/user/repos", bytes.NewReader(jsonBytes))
	if reqErr != nil {
		return nil, reqErr
	}
	request.Header.Set(authHeader, "token "+token)
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 201 {
		return nil, errors.New("github error")
	}
	bodyResp, err := ioutil.ReadAll(response.Body)
	var resp = model.CreateRepoResponse{}
	json.Unmarshal(bodyResp, &resp)
	return &resp, nil
}
