package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var address = "localhost:8080"
var updating = false
var amount = 10000.0
var functionList []string

func main() {
	// Updating server address
	if len(os.Args) > 1 {
		address = os.Args[1] + ":8080"
	}

	link, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	// Start and keep running the replication server
	go listenToReplica()

	// Keep listening for updates and perform them
	// This serial loop makes sure that only one transaction gets hold of the 'updating' variable at a time
	for {
		fmt.Println("Waiting for connection")
		con, errCon := link.Accept()
		if errCon != nil {
			log.Fatal(errCon)
		}
		defer con.Close()
		message, _ := bufio.NewReader(con).ReadString('\n')
		reply := getReplyMessage(message)
		fmt.Fprintln(con, reply+"\n")
	}
}

func listenToReplica() {
	link, err := net.Listen("tcp", strings.TrimRight(address, "8080")+"8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Listening to replica")
		con, errCon := link.Accept()
		if errCon != nil {
			log.Fatal(errCon)
		}
		defer con.Close()
		message, _ := bufio.NewReader(con).ReadString('\n')
		if strings.HasPrefix(message, "SERVER") {
			if !updating {
				fmt.Fprintf(con, fmt.Sprintf("%f\n", amount))
			} else {
				fmt.Fprintf(con, fmt.Sprintf("%f\n", -1.0))
			}
		}
	}
}

func getReplyMessage(msg string) string {
	arr := strings.SplitAfter(msg, " ")
	number := -1.0
	var floatErr error
	command := arr[0]
	if len(arr) > 1 {
		number, floatErr = strconv.ParseFloat(strings.TrimSpace(arr[1]), 64)
		if floatErr != nil {
			log.Fatal(floatErr)
		}
	}

	if strings.HasPrefix(command, "Display") {
		return display()
	} else if strings.HasPrefix(command, "Add") {
		if number < 0 {
			return "Please enter valid add amount"
		}
		if !updating {
			add(number)
			return fmt.Sprintf("Balance after update : %f", amount)
		}
		return "Another transaction in process. Please try again"
	} else if strings.HasPrefix(command, "Withdraw") {
		if number < 0 {
			return "Please enter valid withdraw amount"
		} else if number > amount {
			return "Withdrwal amount is greater than balance"
		}
		if !updating {
			withdraw(number)
			return fmt.Sprintf("Balance after withdrawl : %f", amount)
		}
		return "Another transaction in process. Please try again"
	}
	return "Could not find command. Acceptable values : 'Withdraw', 'Add', 'Display'"
}

func display() string {
	return fmt.Sprintf("Current balance is : %f", amount)
}

func add(aAmount float64) {
	updating = true
	//time.Sleep(2 * time.Second)
	//time.Sleep(10 * time.Second)
	amount += aAmount
	updating = false

}

func withdraw(wAmount float64) {
	updating = true
	amount -= wAmount
	updating = false
}
