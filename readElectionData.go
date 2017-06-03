package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readElectionData(filePath string) Election {
	// Election:
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var election Election
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &election)
	if err != nil {
		fmt.Println(err)
	}
	return election
}
