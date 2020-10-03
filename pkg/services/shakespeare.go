package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type translateResponse struct {
	Contents struct {
		Text        string `json:"text"`
		Translated  string `json:"translated"`
		Translation string `json:"translation"`
	} `json:"contents"`
}

func TranslateToShakespeare(text string) (string, error) {
	payload, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		return "", err
	}

	url := "https://api.funtranslations.com/translate/shakespeare.json"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("shakespeare api response %s: %s", resp.Status, body)
	}

	var result translateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Contents.Translated, nil
}
