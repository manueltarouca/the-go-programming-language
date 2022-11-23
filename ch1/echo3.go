package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Split approach
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	diff := end.Sub(start).Microseconds()
	fmt.Println("Split approach took " + fmt.Sprintf("%d", diff) + " mcs")

	fmt.Println(os.Args[1:])

	fmt.Println(os.Args)

	// Inefficient approach 1
	start = time.Now()
	for i, arg := range os.Args {
		fmt.Println(fmt.Sprintf("%d", i) + " " + arg)
	}
	end = time.Now()
	diff = end.Sub(start).Microseconds()
	fmt.Println("Inefficient approach 1 took " + fmt.Sprintf("%d", diff) + " mcs")

	// Inefficient approach 2
	start = time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end = time.Now()
	diff = end.Sub(start).Microseconds()
	fmt.Println("Inefficient approach 2 took " + fmt.Sprintf("%d", diff) + " mcs")
}
