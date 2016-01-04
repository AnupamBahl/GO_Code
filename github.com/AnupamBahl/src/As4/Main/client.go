package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var clientLamportClock = 0

func registerQuestion(str string, add string) {
	clientLamportClock++
	str = strings.Replace(str, " ", "_", -1)
	resp, err := http.Get("http://" + add + ":8080/ServerHelper?Client=Yes&Clock=" + strconv.Itoa(clientLamportClock) + "&Question=" + str)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Panic(readErr)
	}
	fmt.Println("\nFrom server ::: " + string(data))
	fmt.Printf("Client Time : %v\n\n", time.Now())
}

func main() {
	add := "localhost"
	if len(os.Args) > 1 {
		add = os.Args[1]
	}

	//Take question as std input and call GET to server
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Question : ")
	str, _ := reader.ReadString('\n')
	str = str[0 : len(str)-1]

	//Keep requesting for questions until user types 'exit'
	for !strings.HasPrefix(str, "exit") {
		registerQuestion(str, add)
		fmt.Print("Enter Question : ")
		str, _ = reader.ReadString('\n')
		str = str[0 : len(str)-1]
	}

}
