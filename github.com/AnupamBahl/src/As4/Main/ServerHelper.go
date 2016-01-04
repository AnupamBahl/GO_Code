package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var lamportClockHelper = 0
var shutdown = false

func tryCalling(add string) {
	lamportClockHelper++
	resp, err := http.Get("http://" + add + ":8080/ServerHelper?Client=No&Clock=" + strconv.Itoa(lamportClockHelper))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	//Setting up IP address
	add := "localhost"
	if len(os.Args) > 1 {
		add = os.Args[1]
	}
	for i := 0; i < 20; i++ {
		tryCalling(add)
		fmt.Println("Topic Number : " + strconv.Itoa(lamportClockHelper))
		time.Sleep(1000 * time.Millisecond)
	}

	//set for shutdown
	resp, err := http.Get("http://" + add + ":8080/ServerHelper?Client=Shutdown&Clock=" + strconv.Itoa(lamportClockHelper))
	if err != nil {
		os.Exit(0)
	}
	resp.Body.Close()
}
