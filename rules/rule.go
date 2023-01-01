package rules

import "strconv"

type Rule interface {
	DoesApply(turn string) bool
	Apply(index int, frames []string) int
}

func CountThrow(throw string) int {
	if isStrike(throw) {
		return 10
	}
	convertedNumber, _ := strconv.Atoi(throw)
	return convertedNumber
}

func isStrike(turn string) bool {
	return "x" == turn
}
