package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func jsonReq(url string, reqJSON interface{}) error {
	requestBody, err := json.Marshal(reqJSON)
	if err != nil {
		return err
	}

	url = notaAddr + url

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.Status != "202 Accepted" {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
