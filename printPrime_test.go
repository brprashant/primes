package main

import (
	"reflect"
	"testing"
)


func TestNumberSeries(t *testing.T){
	c := make(chan int)
	go numbers(c)

	val :=2
	for r:=0;r<1000000;r++{
		if ret := <-c; ret!=val {
			t.Errorf("expected %d, got %d", val, ret)
		}
		val =val+1
	}
}

func TestIsPrime(t *testing.T){
	set := []struct {
		ip int
		expected bool
	} {
		{ 1, true, },
		{ 2, true, },
		{ 3, true, },
		{ 4, false, },
		{ 5, true, },
		{ 6, false, },
		{ 7, true, },
		{ 8, false, },
		{ 9, false, },
		{ 10, false, },
		{ 11, true, },
		{ 29, true, },
		{ 523, true, },
		{ 524, false, },
		{ 525, false, },
		{ 526, false, },
		{ 527, false, },
		{ 528, false, },
		{ 529, false, },
		{ 541, true, },
		{ 542, false, },
	}
	for _,sample := range set {
		if val := isPrime(sample.ip); val != sample.expected {
			t.Errorf("Is Prime failed for %d", sample.ip)
		}
	}
}

func TestNPrimes(t *testing.T){
	set := []struct {
		n int
		expected []int
	}{
		{1, []int{2}},
		{2, []int{2,3}},
		{3, []int{2,3,5}},
		{4, []int{2,3,5,7}},
		{5, []int{2,3,5,7,11}},
		{5, nil},
		{10, nil},
		{100, nil},
		{200, nil},
		{500, nil},
		{1000, nil},
		//{2000, nil},
		//{5000, nil},
	}
	for _,sample := range set {
		val := NPrimes(sample.n)
		if len(val) != sample.n {
			t.Errorf("Failed expected length of get N Primes for %d ", sample.n)
		}
		if sample.expected!= nil {
			if !reflect.DeepEqual(val, sample.expected) {
				t.Errorf("expected prime series fails deep equal with response for %d", sample.n)
			}
		}
		for _,v:= range val {
			if !isPrime(v) {
				t.Errorf("A returned value %d is not prime. Found when running n primes for %d", v, sample.n)
			}
		}
	}
}



func TestIsPrimeGivenPrimes(t *testing.T){

	set := []struct {
		series []int
		n int
		expected bool
	}{
		{ []int{}, 2, true },
		{ []int{2}, 3, true },
		{ []int{2,3}, 4, false },
		{ []int{2,3}, 5, true },
		{ []int{2,3,5}, 6, false },
		{ []int{2,3,5}, 7, true },
		{ []int{2,3,5,7}, 8, false },
		{ []int{2,3,5,7}, 9, false },
		{ []int{2,3,5,7}, 10, false },
		{ []int{2,3,5,7}, 11, true },
		{ []int{2,3,5,7,11}, 12, false },
		{ []int{2,3,5,7,11}, 13, true },
	}

	for _,sample := range set {
		if val := givenPrimesIsPrime(sample.series,sample.n); val != sample.expected {
			t.Errorf("Failed IsPrime for %d, with series %v", sample.n, sample.series)
		}
	}

}

func TestRPad(t *testing.T){
	set := []struct{
		num int
		n int
		expected string
		expectedLen int
	}{
		{0, 10, "0         ", 10},
		{123, 10, "123       ", 10},
		{1234, 10, "1234      ", 10},
		{1234567890183, 10, "1234567890183", 13},
	}

	for _,sample := range set {
		val := rPad(sample.num,sample.n)
		if val != sample.expected {
			t.Errorf("RPad Failed for number %d, for n %d. Expected %v, Got %v.", sample.num, sample.n, sample.expected,val)
		}
		if len(val) != sample.expectedLen {
			t.Errorf("RPad Failed for number %d, for n %d. Expected %d, Got %v.", sample.num, sample.n, sample.expectedLen,len(val))
		}
	}

}


// used in test only
// for numbers greater than 1
// check test for uses
func isPrime(n int) bool {
	for i:=2 ; i<n; i++{
		if n%i == 0 {
			return false
		}
	}
	return true
}
