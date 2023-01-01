package rules

import "strconv"

type rule interface {
	DoesApply(turn string) bool
	Apply(score int, index int, frames []string) int
}

func Translate(value string) int {
	if len(value) == 2 { // second throw was a miss, e.g. 7-
		value = value[0:1]
	}
	if len(value) == 3 { // did not pick up the spare, e.g. 7/2
		firstTry := Translate(value[0:1])
		secondTry := Translate(value[2:3])
		return firstTry + secondTry
	}
	if "x" == value {
		return 10
	}
	convertedNumber, _ := strconv.Atoi(value)
	return convertedNumber
}
