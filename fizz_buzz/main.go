package main

import "fmt"

/*

The FizzBuzz problem is a classic test given in coding interviews. The task is simple: Print integers 1 to N, but print “Fizz” if an integer is divisible by 3, “Buzz” if an integer is divisible by 5, and “FizzBuzz” if an integer is divisible by both 3 and 5.

7 Bazz

*/

var Checker = map[int]string{
	3:  "Fizz",
	5:  "Buzz",
	7:  "Bazz",
	11: "Fazz",
}

var N = 1000

func main() {

	for i := 0; i < N; i++ {
		r := FizzBuzzCheck(i)
		if r != "" {
			fmt.Printf("%s: %d\n", r, i)
		}
	}

}

func FizzBuzzCheck(i int) string {
	rString := ""

	for k, v := range Checker {
		if i%k == 0 {
			rString = rString + v
		}
	}

	return rString
}
