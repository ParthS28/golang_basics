package main

import (
	"fmt"
	"strconv"
	"sort"
)


func main() {
	arr := make([]int, 3)
	for i := 0; ; i++ {
		var input string
		fmt.Println("Enter a number to be added or X to quit")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			break
		} else {
			num, er := strconv.Atoi(input)
			if i < 3 {
				if er != nil {
				continue
				}
				arr[2-i] = num
				//sort.Ints(arr)
				//fmt.Println(arr)
				//continue	
			} else {
			if er != nil {
				continue
			}
			arr = append(arr, num)
			}
			sort.Ints(arr)
			fmt.Println(arr)
		}
	}
}
