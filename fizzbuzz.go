// Copyright © 2016 Marko Tunjic.
package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := make(chan int)
	results := make(chan string)

	go generate(numbers, 100)
	go filter(numbers, results)

	printer(results)
}

func printer(in <-chan string) {
	for v := range in {
		fmt.Println(v)
	}
}

func generate(out chan<- int, n int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

func filter(in <-chan int, out chan<- string) {
	for i := range in {

		switch {
		case i%15 == 0:
			out <- "FizzBuzz"
		case i%3 == 0:
			out <- "Fizz"
		case i%5 == 0:
			out <- "Buzz"
		default:
			out <- strconv.Itoa(i)
		}
	}
	close(out)
}
