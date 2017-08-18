package main

import (
	"encoding/json"
	"fmt"
	"github.com/jdextraze/go-gesclient"
	"github.com/jdextraze/go-gesclient/client"
	"github.com/satori/go.uuid"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var (
	errorChan      = make(chan *client.EventData)
	closeErrorChan = make(chan os.Signal, 1)
)

type errorEvent struct {
	Message string `json:"message"`
}

func createErr(errType string, err error) {
	var (
		errStr = fmt.Sprintf("%v", err)
		data   []byte
	)

	log.Printf("%v: %v", errType, err)

	data, _ = json.Marshal(errorEvent{Message: errStr})
	errorChan <- client.NewEventData(uuid.NewV4(), errType, true, data, nil)
}

func errorStream() {
	const stream = "Errors"

	uri, err := url.Parse(gesAddr)
	if err != nil {
		log.Fatalf("Error parsing address: %v", err)
	}

	c, err := gesclient.Create(client.DefaultConnectionSettings, uri, "Publisher")
	if err != nil {
		log.Fatalf("Error creating connection: %v", err)
	}

	if err := c.ConnectAsync().Wait(); err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	signal.Notify(closeErrorChan, os.Interrupt)

	for {
		select {
		case errEventData := <-errorChan:
			result := &client.WriteResult{}
			task, err := c.AppendToStreamAsync(stream, client.ExpectedVersion_Any, []*client.EventData{errEventData}, nil)
			if err != nil {
				log.Printf("Error occured while appending to error stream: %v", err)
			} else if err := task.Result(result); err != nil {
				log.Printf("Error occured while waiting for result of appending to error stream: %v", err)
			}

		case <-closeErrorChan:
			c.Close()
			time.Sleep(10 * time.Millisecond)
			return
		}
	}
}
