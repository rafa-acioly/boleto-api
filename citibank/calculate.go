package citibank

import (
	"strconv"
	"time"
)

func mod11(valueSequence string) string  {
	digit := 0
	sum := 0
	weight := 2

	var values []int

	for _, r := range valueSequence {
		c := string(r)
		n, _ := strconv.Atoi(c)
		values = append(values, n)
	}

	for  i := len(values)-1; i >= 0; i-- {
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

func mod11Base9(valueSequence string) string {
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
		if weight == 9 {
			weight = 2
		} else {
			weight = weight + 1
		}
	}

	digit = 11 - (sum % 11)

	if digit > 9 || digit == 0 || digit == 1 {
		digit = 1
	}

	return strconv.Itoa(digit)
}


func mod10(valueSequence string) string {
	digit := 0
	sum := 0
	weight := 2
	rest := 0

	var values []int

	for _, r := range valueSequence {
		c := string(r)
		n, _ := strconv.Atoi(c)
		values = append(values, n)
	}

	for i := len(values)-1; i >= 0; i-- {
		rest = values[i] * weight

		if rest > 9 {
			rest = (rest / 10) + (rest % 10)
		}
		sum += rest
		if weight == 2 {
			weight = 1
		} else {
			weight = weight + 1
		}
	}

	digit = (10 - (sum % 10)) % 10

	return strconv.Itoa(digit)
}

func dateDueFactor(dateDue time.Time) string {
	var dateDueFixed = time.Date(1997, 10, 7, 0, 0, 0, 0, time.UTC)
	dif := dateDue.Sub(dateDueFixed)
	factor := int(dif.Hours() / 24)
	if factor <= 0 {
		panic("DateDue must be in the future")
	}
	return strconv.Itoa(factor)
}