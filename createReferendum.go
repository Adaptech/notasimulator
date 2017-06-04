package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func createReferendum(ID string, organizationID string, name string, proposal string, options []string) {

	url := "http://localhost:3001/api/v1/organization/referendum/create"

	optionsJSON, err := json.Marshal(&options)
	if err != nil {
		fmt.Println(err)
	}
	requestBody := fmt.Sprintf("{\n   \"referendumId\": \"%v\",\n   \"organizationId\": \"%v\",\n   \"name\": \"%v\",\n   \"proposal\": \"%v\",\n   \"options\": %v\n}", ID, organizationID, name, proposal, string(optionsJSON))

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
