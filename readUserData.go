package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readUserData(filePath string) (user []User, err error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	var users []User

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
