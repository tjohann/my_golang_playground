// printf via go-routines
package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Println(s)
}


func main() {

	print("\nshort text");
	print("1")
	go print("2")
	// this will take longer then the output of print("3")
	print("3")
	print("\n");
	
	print("\nlong text");
	print("count one")
	// this will take longer then the output of print("3")
	go print("count two") 
	print("count three")
	print("\n");

	// this is needed, because go wont wait for go-routines
	time.Sleep(1 * time.Second)
}
