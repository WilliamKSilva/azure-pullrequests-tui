package ui

import (
	"encoding/json"
	"fmt"

	utils "github.com/WilliamKSilva/azure-pullrequests-cli/utils"
)

type Repository struct {
    Name string `json:"name"`
}

type PullRequestsData struct {
  Repository Repository `json:"repository"`
  Url string `json:"url"`
  Status string `json:"status"`
  Title string `json:"title"`
}

type PullRequests struct {
  Count int `json:"count"`
  Value []PullRequestsData `json:"value"`
}

func getPullRequests(token string, organization string, project string) (*PullRequests, error) {
    var pullRequests PullRequests

    url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/pullrequests?api-version=7.0", organization, project)  
    patToken := fmt.Sprintf(":%s", token)
    body, err := utils.GetRequest(url, patToken) 

    if err != nil {
        return nil, err
    }

    json.Unmarshal([]byte(*body), &pullRequests)

    return &pullRequests, nil 
}

