package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Value struct {
  Url string `json:"url"`
  Status string `json:"status"`
  Title string `json:"title"`
}

type ResponseBody struct {
  Count int `json:"count"`
  Value []Value `json:"value"`
}

var client = &http.Client{}

func encodeToBase64(value string) string {
	decoded := base64.StdEncoding.EncodeToString([]byte(value))

	return decoded
}

func parseJSON(res *http.Response) string {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func getRequest(url string, token string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err 
	}

	encodedToken := encodeToBase64(token)
	authorizationString := fmt.Sprintf("Basic %s", encodedToken)
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)

	if err != nil {
		return "", err 
	}

	body := parseJSON(res)

	return body, nil 
}

func FetchProjects(token string, organization string) (*ResponseBody, error) {
    url := fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=7.0", organization)  
    body, err := getRequest(url, token) 

    if err != nil {
        return nil, err
    }

    var projects ResponseBody 

    json.Unmarshal([]byte(body), &projects)

    return &projects, nil
}