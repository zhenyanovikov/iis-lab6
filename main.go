package main

import (
	"encoding/json"
	"fmt"
	"github.com/telkomdev/go-stash"
	"os"
	"strconv"
	"time"
)

var logger *stash.Stash
var host string = "localhost"
var port uint64 = 50000

type Message struct {
	Data string `json:"data"`
}

type Log struct {
	Action  string    `json:"action"`
	Time    time.Time `json:"time"`
	Message Message   `json:"message"`
}

func main() {
	logger = initLogger(host, port)

	defer func() {
		logger.Close()
	}()

	var logData Log

	for i := 0; ; i++ {
		logData = Log{
			Action: "get_me",
			Time:   time.Now(),
			Message: Message{
				Data: "get me for me index " + strconv.Itoa(i),
			},
		}

		log(logData)
		logData = Log{
			Action: "register",
			Time:   time.Now(),
			Message: Message{
				Data: "oh wow register on index " + strconv.Itoa(i),
			},
		}

		log(logData)

		fmt.Println("log data: ", logData)

		<-time.After(time.Second * 1)
	}
}

func log(data any) {
	logDataJSON, _ := json.Marshal(data)

	_, err := logger.Write(logDataJSON)
	if err != nil {
		panic(err)
		return
	}
}

func initLogger(host string, port uint64) *stash.Stash {
	var err error
	logger, err = stash.Connect(host, port)
	if err != nil {
		os.Exit(1)
	}
	return logger
}
