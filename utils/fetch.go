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
    if (url == "" || token == "") {
        return nil, errors.New("missing arguments")
    }
    
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err 
	}

	encodedToken := encodeToBase64(token)
	authorizationString := fmt.Sprintf("Basic %s", encodedToken)
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)

	if err != nil {
		return nil, err 
	}

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
