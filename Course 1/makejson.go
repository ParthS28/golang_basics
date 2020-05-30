package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	db := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter a name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Now enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	db["address"] = address
	db["name"] = name

	barr, err := json.Marshal(db)


	if err == nil {
		fmt.Println("JSON file: \n ", string(barr)) 
	} else {
		fmt.Println(err)
	}


}
