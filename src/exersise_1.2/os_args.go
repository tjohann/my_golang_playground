// os.args example
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "

	fmt.Println("programm: " + os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		fmt.Printf("os.Args[%d]: %s\n", i, os.Args[i])
	}

	fmt.Println("kdo-args: " + s)
}
