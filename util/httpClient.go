package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAsType[T any](url string, headers map[string]string) (T, error) {
	body, err := makeGetCall(url, headers)

	if err != nil {
		return *new(T), err
	}

	// Unmarshal the JSON response into the struct
	var response T
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return *new(T), err
	}
	return response, nil
}

func makeGetCall(url string, headers map[string]string) ([]byte, error) {
	fmt.Println("Making GET call to:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Send the request using the client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	return body, nil
}
