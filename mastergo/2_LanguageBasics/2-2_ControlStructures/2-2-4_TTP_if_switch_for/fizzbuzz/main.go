package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {

	for i := 1; i <= n; i++ {
		output := ""
		if i%3 == 0 {
			output = output + "Fizz"
		}
		if i%5 == 0 {
			output = output + "Buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		println(output)
	}
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
