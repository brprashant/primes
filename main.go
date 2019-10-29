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
	fmt.Println("Hello, world.")

	n := flag.Int("n", 10, "suggesting a reasonable int less <5000")
	flag.Parse()

	if *n > 5000 {
		fmt.Println("suggest an input less than 5000")
		return
	}

	primes := nPrimes(*n)

	bigNum := primes[*n-1] * primes[*n-1]
	strBigNum := strconv.Itoa(bigNum)
	maxPad := len(strBigNum)
	space := []byte(" ")

	for _,x := range primes {
		for _,y := range primes {
			num := x*y
			hold := []byte(strconv.Itoa(num))

			for ;maxPad-len(hold)!=0;{
				hold = append(hold, space[0])
			}

			fmt.Print(string(hold) + " ")
		}
		fmt.Printf("\n")
	}

}

func nPrimes(n int) []int{
	var allPrimes []int


	for i:=1;len(allPrimes)<n+1;i++ {
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