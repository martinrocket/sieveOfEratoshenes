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

	var x, y []int

	//for i := 2; i <= 100; i++ {
	for i := 2; i <= 1000; i++ {
		x = append(x, i)

	}

	//fmt.Println(x)

	for j := len(x); j > 0; j-- {
		for i := 0; i < j; i++ {
			//fmt.Printf("%d / %d = %d %d \n", x[j-1], x[i], x[j-1]/x[i], x[j-1]%x[i])
			if x[j-1]%x[i] == 0 && x[j-1]/x[i] != 1 {
				//fmt.Printf(" break %d \n", x[j-1])
				break
			} else {
				if x[j-1]/x[i] == 1 {
					y = append(y, x[j-1])
					break
				}
			}
		}
	}
	slices.Reverse(y)
	fmt.Println(y)

}
