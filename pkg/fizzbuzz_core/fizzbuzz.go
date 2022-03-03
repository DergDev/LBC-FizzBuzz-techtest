package fizzbuzz_core

import "strconv"

//Fizz Buzz Core Logic
func ComputeFizzBuzz(int1 uint8, int2 uint8, limit uint8, str1 string, str2 string) []string {

	fizzBuzz := []string{}
	for i := uint8(1); i <= limit; i++ {
		current := ""
		if i%int1 == 0 {
			current = str1
		}
		if i%int2 == 0 {
			current += str2
		}
		if len(current) == 0 {
			current = strconv.Itoa(int(i))
		}
		fizzBuzz = append(fizzBuzz, current)
	}
	return fizzBuzz
}
