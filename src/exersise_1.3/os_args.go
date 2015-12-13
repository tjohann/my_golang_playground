// os.args example
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", " "

	//fmt.Println("programm: " + os.Args[0])

	fmt.Println("\n\nFirst version -> lang==go")
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		//	fmt.Printf("os.Args[%d]: %s\n", i, os.Args[i])
	}
	//end := time.Now()
	sec := time.Since(start).Seconds()
	//fmt.Println("kdo-args: " + s)
	//fmt.Printf("Start-time: %s \n", start)
	//fmt.Printf("Stop-time: %s \n", end)
	fmt.Printf("Result in sec: %fs \n", sec)
	// Result -> Diff in sec: 0.000040s

	fmt.Println("Second version -> lang==go")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	//end = time.Now()
	sec = time.Since(start).Seconds()
	//fmt.Println("kdo-args: " + s)
	//fmt.Printf("Start-time: %s \n", start)
	//fmt.Printf("Stop-time: %s \n", end)
	fmt.Printf("Result in sec: %fs \n", sec)
	// Result -> Diff in sec: 0.000019s
}
