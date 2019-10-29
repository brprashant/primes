package main

import (
	"fmt"
	"strconv"
)


//Write a program that prints out a multiplication table of the first 10 prime numbers.
//	The program must run from the command line and print one table to STDOUT.
//	The first row and column of the table should have 10 primes, with each cell containing the
//product of the primes for the corresponding row and column.
func main() {
	fmt.Println("Hello, world.")

	// accept numbers
	// print table of n_prime * n_prime
	// first row and column is primes

	// GENERATE primes []int
	primes := nPrimes(10)

	// primes := []int
	// for every number in primes
		// print current prime * every number in prime and print
		// new line
	// end
	for _,x := range primes {
		for _,y := range primes {
			fmt.Printf(strconv.Itoa(x*y) + " ")
		}
		fmt.Printf("\n")
	}

}

func nPrimes(n int) []int{
	var allPrimes []int


	for i:=1;len(allPrimes)<n;i++ {
		num := i
		if isPrime(num) {
			allPrimes = append(allPrimes, num)
		}
	}
	return allPrimes
}

func isPrime(n int) bool {
	for i:=2 ; i<n; i++{
		if n%i == 0 {
			return false
		}
	}
	return true
}