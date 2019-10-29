package main

import (
	"flag"
	"fmt"
	"strconv"
)


//Write a program that prints out a multiplication table of the first 10 prime numbers.
//	The program must run from the command line and print one table to STDOUT.
//	The first row and column of the table should have 10 primes, with each cell containing the
//product of the primes for the corresponding row and column.
func main() {
	n := flag.Int("n", 10, "suggesting a reasonable int less <5000")
	flag.Parse()

	if *n > 5000 || *n<1 {
		fmt.Println("suggest a positive input between 1 and 5000")
		return
	}

	// primes including 1!
	primes := append([]int{1}, NPrimes(*n)...)

	bigNum := primes[*n-1] * primes[*n-1]
	strBigNum := strconv.Itoa(bigNum)
	maxPad := len(strBigNum)

	for _,x := range primes {
		for _,y := range primes {
			num := x*y
			fmt.Print(rPad(num, maxPad+1))
		}
		fmt.Printf("\n")
	}

	return
}

// number generating channel :shrug:
func numbers(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Given n, returns a slice of first n primes.
func NPrimes(n int) []int{
	var allPrimes []int

	c := make(chan int)
	go numbers(c)

	for ;len(allPrimes)<n; {
		num := <-c
		if givenPrimesIsPrime(allPrimes, num) {
			allPrimes = append(allPrimes, num)
		}
	}
	return allPrimes
}

// locally used in generating n primes.
// given all known primes till n, and an n, finds if n is divisble by any known primes so far.
// used to find if a number is prime, quicker than checking all numbers till n.
func givenPrimesIsPrime(known []int, n int) bool {
	for _,v := range known {
		if n%v == 0 {
			return false
		}
	}
	return true
}

// given int, fills ' ' spaces to the right and returns string.
// given number, and an expected padded length
// returns a string of with at least provided length.
// Does not expect paddedLength to be less than length of num
// Returns number as string if paddedLength is lesser than number provided.
func rPad(number int, paddedLength int) string {
	hold := []byte(strconv.Itoa(number))
	space := []byte(" ")
	if len(hold)> paddedLength {
		return string(hold)
	}

	for ; paddedLength-len(hold)!=0;{
		hold = append(hold, space[0])
	}

	return string(hold)
}
