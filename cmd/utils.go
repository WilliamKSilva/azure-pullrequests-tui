package cmd

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

var client = &http.Client{}

func encodeToBase64(value string) string {
	decoded := base64.StdEncoding.EncodeToString([]byte(value))

	return decoded
}

func GetRequest(url string, token string) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	authorizationString := fmt.Sprintf("Basic %s", token)

	authorization := encodeToBase64(authorizationString)

	req.Header.Add("Authorization", authorization)

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
