package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readElectionData(filePath string) (*Election, error) {
	var election Election

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &election)
	if err != nil {
		return nil, err
	}

	return &election, nil
}
