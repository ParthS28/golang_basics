package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    fmt.Println("Please input numbers(less than 10, separated by space):")
    br := bufio.NewReader(os.Stdin)
    a, _, _ := br.ReadLine()
    ns := strings.Split(string(a), " ")
    var arr []int
    for _, s := range(ns) {
      n, _ := strconv.Atoi(s)
      arr = append(arr, n)
    }
    Bubble(arr)
    fmt.Println(arr)
}

func Bubble(a []int) {    
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a)-1-i; j++  {
            if a[j+1] < a[j] {
                Swap(a,j)
            }
        }
    }
}

func Swap(a []int, i int) {
    var temp int
    temp = a[i]
    a[i] = a[i+1]
    a[i+1] = temp
}