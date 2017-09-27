package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	Slack struct {
		Channel  string `json:"channel"`
		Username string `json:"username"`
		IconUrl  string `json:"icon_url"`
		Text     string `json:"text"`
	}
)

// RequestPost
// JSONをしているのURLにPOSTする
func RequestPost(json []uint8, url string) (rst []byte, err error) {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(json)),
	)
	if err != nil {
		log.Printf("data : %v", err)
		return rst, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("data : %v", err)
		return rst, err
	}
	defer res.Body.Close()

	rst, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("data : %v", err)
		return rst, err
	}
	return rst, err
}

// SendSlack
// webhookを使用してSLACKに通知を送る
func SendSlack(slack Slack, slackUrl string) error {
	json, err := json.Marshal(slack)
	if err != nil {
		log.Printf("data : %v", err)
		return err
	}

	_, err = RequestPost(json, slackUrl)
	return err
}
