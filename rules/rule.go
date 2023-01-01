package rules

import "strconv"

type Rule interface {
	DoesApply(turn string) bool
	Apply(index int, frames []string) int
}

func Translate(turn string) int {
	if len(turn) == 2 && turn[1] == '-' { // second throw was a miss, e.g. 7-
		turn = turn[0:1]
	}
	if len(turn) == 3 { // did not pick up the spare, e.g. 7/2
		firstTry := Translate(turn[0:1])
		secondTry := Translate(turn[2:3])
		return firstTry + secondTry
	}
	if "x" == turn {
		return 10
	}
	convertedNumber, _ := strconv.Atoi(turn)
	return convertedNumber
}
