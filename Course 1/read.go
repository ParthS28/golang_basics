package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	first_name string
	last_name string
}

func trunc(n string) string {
	if len(n) > 20 {
		return n[0:20]
	}
	return n
}

func main() {
	fmt.Println("Enter name of .txt file to read:")
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		fmt.Println("Cannot read the file", input.Err())
		return
	}
	fileName := input.Text()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR: ", err, "Failed to open filename: ", fileName)
		file.Close()
		return
	}

	reader := bufio.NewScanner(file)
	names := make([]name, 0, 3)
	for reader.Scan() {
		e := strings.Split(reader.Text(), " ")
		names = append(names, name{first_name: trunc(e[0]),
			last_name: trunc(e[1])})
	}
	file.Close()

	fmt.Printf("%-20s \t %-20s\n", "First Name", "Last Name")
	for _, n := range names {
		fmt.Printf("%-20s \t %-20s\n", n.first_name, n.last_name)
	}
}
