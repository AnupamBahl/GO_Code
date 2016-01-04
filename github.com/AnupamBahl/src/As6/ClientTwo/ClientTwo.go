package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var serverAddress = "localhost:8040"

func main() {
	if len(os.Args) > 1 {
		serverAddress = os.Args[1] + ":8040"
	}

	for {
		// Create connection
		clientCon, err := net.Dial("tcp", serverAddress)
		if err != nil {
			log.Fatal(err)
		}
		// Take command line input and send to server as command
		fmt.Println("Waiting for input")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.HasPrefix(command, "exit") {
			clientCon.Close()
			return
		}
		fmt.Fprintf(clientCon, command+"\n")
		//Listen to reply
		reply, _ := bufio.NewReader(clientCon).ReadString('\n')
		fmt.Println(reply)
		clientCon.Close()
	}
}
