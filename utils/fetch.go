package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

func GetRequest(url string, token string) (*string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err 
	}

	encodedToken := encodeToBase64(token)
	authorizationString := fmt.Sprintf("Basic %s", encodedToken)
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)

    if res.StatusCode == 203 {
        return nil, errors.New("Error: Invalid PAT (Personal Access Token)")
    }

    if res.StatusCode == 404 {
        return nil, errors.New("Error: Azure DevOps organization name not found")
    }

	if err != nil {
		return nil, err 
	}

    defer res.Body.Close()
	body := parseJSON(res)

	return &body, nil 
}

func encodeToBase64(value string) string {
    encoded := base64.StdEncoding.EncodeToString([]byte(value))

    return encoded 
}

func parseJSON(res *http.Response) string {
    body, err := io.ReadAll(res.Body)

    if err != nil {
        panic(err)
    }

    return string(body)
}
