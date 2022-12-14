package cmd

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

func EncodeToBase64(value string) string {
	decoded := base64.StdEncoding.EncodeToString([]byte(value))

	return decoded
}

func ParseJSON(res *http.Response) string {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func GetRequest(url string, token string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err 
	}

	encodedToken := EncodeToBase64(token)
	authorizationString := fmt.Sprintf("Basic %s", encodedToken)
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)

	if err != nil {
		return "", err 
	}

	body := ParseJSON(res)

	return body, nil 
}
