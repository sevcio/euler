/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

var primeNumbers []int

func main() {

	largeNumber := 600851475143
	for i := 2; ; i++ {
		if largeNumber%i == 0 {
			fmt.Printf("Prime numbers is %d of number %d and equals to %d\n", i, largeNumber, largeNumber/i)
			primeNumbers = append(primeNumbers, i)
			largeNumber = largeNumber / i
			i = 2
		}

		if largeNumber <= 1 {
			break
		}
	}
}
