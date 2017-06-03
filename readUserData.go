package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readUserData(filePath string) []User {
	// Users:
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var users []User
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
