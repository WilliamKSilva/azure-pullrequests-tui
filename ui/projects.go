package ui

import (
	"encoding/json"
	"fmt"

   utils "github.com/WilliamKSilva/azure-pullrequests-cli/utils"
)

type Value struct {
  Url string `json:"url"`
  Description string `json:"description"`
  Status string `json:"status"`
  Name string `json:"name"`
}

type Projects struct {
  Count int `json:"count"`
  Value []Value `json:"value"`
}

func getProjects(token string, organization string) (*Projects, error) {
    var projects Projects
    url := fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=7.0", organization)  
    patToken := fmt.Sprintf(":%s", token)

    body, err := utils.GetRequest(url, patToken) 

    if err != nil {
        return nil, err 
    }

    json.Unmarshal([]byte(*body), &projects)

    return &projects, nil 
}
