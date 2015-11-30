// os.args example
package main

import (
	"fmt"
	"os"
)

func main() {

	s, sep := "", " "

	for _, arg := range os.Args[1:] {
		s += arg + sep
	}

	fmt.Println("programm: " + os.Args[0])
	fmt.Println("kdo-args: " + s)
}
