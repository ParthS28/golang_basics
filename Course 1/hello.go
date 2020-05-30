package main
import (
    "fmt"
    "strconv"
)

func main() {
    // Declaring the input string 
    var input string
    // Prompt the user for input of floating point number
    fmt.Printf("Please input a float point number:")
    // Reading the input to a string 
    fmt.Scan(&input)
    // Convert to a float
    f, err := strconv.ParseFloat(input, 64)
    // If able to convert, print, else prompt the user for wrong input
    if err == nil {
            i := int(f) // Convert to int to truncate
	    fmt.Printf("The truncated float is: %d ",i) //print the int
    } else {
	    fmt.Printf("The input is not a float")
    }
}