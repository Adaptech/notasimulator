package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func castVote(referendumID string, voterID string, vote string) {

	url := "http://localhost:3001/api/v1/organization/referendum/vote"

	requestBody := fmt.Sprintf("{\n   \"referendumId\": \"%v\",\n   \"voterId\": \"%v\",\n   \"vote\": \"%v\"\n}", referendumID, voterID, vote)
	payload := strings.NewReader(requestBody)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.Status != "202 Accepted" {
		fmt.Println(res)
		fmt.Println(string(body))
	}
}
