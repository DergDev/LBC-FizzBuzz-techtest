package fizzbuzz_core

import (
	"reflect"
	"testing"
)

func TestComputeFizzBuzz(t *testing.T) {
	t.Run("Optimal use case - int1: 5 int1: 4 limit:20 str1:fizz str2:buzz", func(t *testing.T) {
		got := ComputeFizzBuzz(uint8(5), uint8(4), uint8(20), "fizz", "buzz")
		want := []string{
			"1", "2", "3", "buzz", "fizz", "6", "7", "buzz", "9", "fizz", "11", "buzz", "13", "14", "fizz", "buzz", "17", "18", "19", "fizzbuzz",
		}

		assertMessage(t, got, want)
	})

	t.Run("check multiple of 1", func(t *testing.T) {
		got := ComputeFizzBuzz(1, 2, 5, "Foo", "Bar")
		want := []string{"Foo", "FooBar", "Foo", "FooBar", "Foo"}

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got []string, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s but want %s", got, want)
	}
}
