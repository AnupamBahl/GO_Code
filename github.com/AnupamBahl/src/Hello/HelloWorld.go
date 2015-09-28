package main

import (
	"fmt"
	"sort"
)

func getrange() {
	dataMap := make(map[int]string)
	dataMap[0] = "ldkjs"
	dataMap[1] = "32kejd"
	var keys []int
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := range keys {
		fmt.Println(dataMap[i])
	}
}

func main() {
	/*	str := ""
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Hello World")
		fmt.Println(rand.Intn(5) + 5)
		fmt.Print(strconv.Itoa(time.Now().Second()) + ".")
		fmt.Println(time.Now())
		time.Sleep(500 * time.Millisecond)
		str = fmt.Sprintf("Update ::: Hi Helper. Lamport clock :.\tTimestamp : %s", time.Now())
		fmt.Println(str)*/
	getrange()
}
