package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var lamportClock = 0
var dataMap = make(map[int]string)
var lamportNeed = false
var counter = 0

//Lamport implementation for helper as well
func updateClock(clock int) {
	if clock > lamportClock {
		lamportClock = clock
		lamportNeed = true
	} else {
		lamportClock++
	}
}

//Getting response strings
func getHelperStr() string {
	var str = ""
	if lamportNeed {
		str = fmt.Sprintf("Update ::: Helper Helped. Lamport clock : %d.\tTimestamp : %v", lamportClock, time.Now())
	} else {
		str = fmt.Sprintf("Increment ::: Helper Helped. Lamport clock : %d.\tTimestamp : %v", lamportClock, time.Now())
	}
	return str
}
func getClientStr(question string) string {
	var str = ""
	if lamportNeed {
		str = fmt.Sprintf("Update ::: Request recorded at internal clock : %d\t.Timestamp : %v", lamportClock, time.Now())
	} else {
		str = fmt.Sprintf("Increment ::: Request recorded at internal clock : %d\t.Timestamp : %v", lamportClock, time.Now())
	}
	return str
}

//Function to create final logs
func logData() {
	//Ordering data required
	var keys []int
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	file, fileErr := os.Create("Log.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	for i := range keys {
		file.Write([]byte(dataMap[i]))
	}
	file.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
	//Get data from url
	var client = r.URL.Query().Get("Client")
	var clockStr = r.URL.Query().Get("Clock")
	var str = ""
	counter++
	clock, err := strconv.Atoi(clockStr)
	if err != nil {
		log.Fatal(err)
	}
	//update clock
	updateClock(clock)

	if client == "No" {
		str = getHelperStr()
		dataMap[counter] = str + "\n"
		fmt.Fprintf(w, str)
	} else if client == "Shutdown" {
		//On stop, create log file and save data in it.
		logData()
		os.Exit(0)
	} else if client == "Yes" {
		question := r.URL.Query().Get("Question")
		question = strings.Replace(question, "_", " ", -1)
		str = getClientStr(question)
		dataMap[counter] = "\n_____________________________\n" + str + "\n" + "Question ::: " + question + "_____________________________\n"
		fmt.Fprintf(w, "Recorded Question : "+question)
	}

	lamportNeed = false
}

func main() {
	add := ""
	if len(os.Args) > 1 {
		add = os.Args[1]
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(add+":8080", nil)
}
