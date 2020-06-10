package main

import (
	"fmt"
	"time"
)
// Race condition is when multiple threads are trying to access and manipulate the same variable or I/O

func fn (x string){
	fmt.Println(x)
}

func main(){
	for i:=0;i<3;i++ {
		go fn ("Hello")
		go fn ("World")
		time.Sleep(1000 * time.Millisecond)
	}
}
/* 
Here, we have a function that prints the string passed to it.
Each time you run it you can see, the output is different, which proves the race condition and interleaving occuring. 

// func add_one(pt *int) {
//   (*pt)++
//   fmt.Println(*pt)
// }

// func sub_one(pt *int) {
//   (*pt)--
//   fmt.Println(*pt)
// }

// func main() {
//   i := 0
//   go add_one(&i)    // This function gets called as a go routine and then the control follows through
//   // time.Sleep(100 * time.Millisecond)  // If you uncomment this line, you can see that even 1 is printed
//   sub_one(&i)		// This function completes execution before the goroutine and so we can see -1 as the first output
//   add_one(&i)		// After sub_one() control falls through to this function and completes add_one()
//   i++
//   fmt.Println(i)

// }

// // The whole program completes before the go routine and so we don't see the result of the goroutine