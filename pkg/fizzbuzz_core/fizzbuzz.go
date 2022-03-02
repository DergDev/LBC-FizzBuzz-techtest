package fizzbuzz_core

import "strconv"

//Generate Fizz Buzz from the different
func ComputeFizzBuzz(int1 uint8, int2 uint8, limit uint8, fizz string, buzz string) []string {

	fizzBuzz := []string{}
	for i := uint8(1); i < limit; i++ {
		current := ""
		if i%int1 == 0 {
			current = "Fizz"
		}
		if i%int2 == 0 {
			current += "Buzz"
		}
		if len(current) == 0 {
			current = strconv.Itoa(int(i))
		}
		fizzBuzz = append(fizzBuzz, current)
	}
	return fizzBuzz
}
