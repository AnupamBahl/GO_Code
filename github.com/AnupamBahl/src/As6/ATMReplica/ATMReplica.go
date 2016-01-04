package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var amount = 10000.0
var serverAddress = "localhost:8000"

func main() {
	//Set up main server address
	if len(os.Args) > 1 {
		serverAddress = os.Args[1] + ":8000"
	}
	//Listen client requests. Only data display is entertained at the moment
	link, err := net.Listen("tcp", "localhost:8040")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Waiting for connection")
		con, errCon := link.Accept()
		if errCon != nil {
			log.Fatal(errCon)
		}
		defer con.Close()
		message, _ := bufio.NewReader(con).ReadString('\n')
		str := displayData(strings.TrimSpace(message))
		fmt.Fprintf(con, str+"\n")
	}

}

func displayData(str string) string {
	if strings.HasPrefix(str, "Display") {
		if confirmData() {
			return fmt.Sprintf("Current Balance is : %f", amount)
		}
	}
	return "Data is being updated. Try again later"
}

func confirmData() bool {
	// Create connection
	for i := 0; i < 3; i++ {
		replicaCon, err := net.Dial("tcp", serverAddress)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Server is Busy, try Again later")
		}
		fmt.Fprintf(replicaCon, "SERVER\n")
		reply, _ := bufio.NewReader(replicaCon).ReadString('\n')
		amount, _ = strconv.ParseFloat(strings.TrimSpace(reply), 64)
		if amount > 0 {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}
