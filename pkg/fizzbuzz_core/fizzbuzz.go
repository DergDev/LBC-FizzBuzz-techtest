package fizzbuzz_core

import (
	"strconv"
)

//Fizz Buzz Core Logic
func ComputeFizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) []string {

	fizzBuzz := []string{}
	for i := int(1); i <= limit; i++ {
		current := ""
		if i%int1 == 0 {
			current = str1
		}
		if i%int2 == 0 {
			current += str2
		}
		if len(current) == 0 {
			current = strconv.Itoa(i)
		}
		fizzBuzz = append(fizzBuzz, current)
	}
	return fizzBuzz
}
