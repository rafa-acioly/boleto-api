package util

import (
	"strconv"
)

//Mod11 calculate Mod11 DV from string
func Mod11(valueSequence string) string {
	digit := 0
	sum := 0
	weight := 2

	var values []int

	for _, r := range valueSequence {
		c := string(r)
		n, _ := strconv.Atoi(c)
		values = append(values, n)
	}

	for i := len(values) - 1; i >= 0; i-- {
		sum += values[i] * weight

		if weight < 9 {
			weight = weight + 1
		} else {
			weight = 2
		}
	}

	digit = 11 - (sum % 11)

	if digit > 9 {
		digit = 0
	}

	return strconv.Itoa(digit)
}
