package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func registerVoter(ID string, organizationID string, firstname string, lastname string, streetAddress string, postOfficeBoxNumber string, addressLocality string, addressRegion string, postalCode string, addressCountry string) {

	url := "http://localhost:3001/api/v1/organization/voter/register"

	requestBody := fmt.Sprintf("{\n   \"voterId\": \"%v\",\n   \"organizationId\": \"%v\",\n   \"firstname\": \"%v\",\n   \"lastname\": \"%v\",\n   \"address\": {\n      \"streetAddress\": \"%v\",\n      \"postOfficeBoxNumber\": \"%v\",\n      \"addressLocality\": \"%v\",\n      \"addressRegion\": \"%v\",\n      \"postalCode\": \"%v\",\n      \"addressCountry\": \"%v\"\n   }\n}", ID, organizationID, firstname, lastname, streetAddress, postOfficeBoxNumber, addressLocality, addressRegion, postalCode, addressCountry)
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
