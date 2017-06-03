package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func createElectionAdmin(id string) {

	url := "http://localhost:3001/api/v1/electionadmin/create"

	requestBody := fmt.Sprintf("{\n   \"electionAdminId\": \"%v\",\n   \"firstname\": \"Joe\",\n   \"lastname\": \"Admin\",\n   \"address\": {\n      \"streetAddress\": \"405 E. Stueben\",\n      \"postOfficeBoxNumber\": null,\n      \"addressLocality\": \"Bingen\",\n      \"addressRegion\": \"WA\",\n      \"postalCode\": \"98605\",\n      \"addressCountry\": \"US\"\n   }\n}", id)
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
