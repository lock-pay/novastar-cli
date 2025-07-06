package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Request(payloadData any, url, method, token string) (json.RawMessage, error) {
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadBytes)))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", token)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	responseWrapper := ResponseWrapper{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return nil, err
	}

	if responseWrapper.Code != 0 {
		return nil, fmt.Errorf("Failed Request: %s", responseWrapper.Message)
	}
	return responseWrapper.Data, nil
}

func UploadFile(url, token string, fileContent io.Reader, targetPath string) (json.RawMessage, error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, fileContent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/octet-stream")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return json.RawMessage(body), nil
}
