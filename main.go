package main

/* The Sieve of Eratosthenes is an ancient algorithm for finding
all prime numbers up to a specified integer. It works by
iteratively marking the multiples of each prime number as composite,
starting from 2, until reaching the square root of the specified limit. */

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("\n")
	fmt.Println("Sieve of Eratoshenes\n")

	//Create a slice with 100 integers

	var x []int

	//for i := 2; i <= 100; i++ {
	for i := 2; i <= 10; i++ {
		x = append(x, i)

	}

	fmt.Println(x)
	sliceLength := len(x)
	for i := 0; i < sliceLength-1; {
		x = slices.Delete(x, len(x)-1, len(x))
		sliceLength = len(x)
		fmt.Println(x)

	}
}
