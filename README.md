# primes

Problem statment 

Objective
Write a program that prints out a multiplication table of the first 10 prime numbers.
The program must run from the command line and print one table to STDOUT.
The first row and column of the table should have 10 primes, with each cell containing the
product of the primes for the corresponding row and column.
Notes
● Consider complexity. How fast does your code run? How does it scale?
● Consider cases where we want more than N primes.
● Do not use the Prime class from stdlib (write your own code).
● Write tests. Demonstrate TDD/BDD.


>Suggested to be written in Python or Ruby.
>>Chose `go`, because of current usage familiarity. 

Current Status : 
Reuires GO to be installed for clean build and run.
~Provided executable will work on a mac. But not tested or copiled for external machines.~
Compiles program. 
Runs for default value of 10.
Can accept positive integers from 1 to 5000.
First row and column are prime numbers.
First Row's first column is `1`.
Acknoledge it is not a prime. But still prints `1`.
Given n, prints n+1 rows.


# Instructions 
1. Install GO in your computer. 
2. Pull this repo.
3. cd into this project folder.
4. $`go run printPrimes.go` for default 10 value output
5. $'go run printPrimes.go -n <number>' will print for n numbers. 

