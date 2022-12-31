package bowling

import "strconv"

func Score(frames []string) int {
	score := 0
	for _, turn := range frames {
		score += translate(turn)
	}
	return score
}

func translate(value string) int {
	if len(value) == 2 { // second throw was a miss, e.g. 7-
		value = value[0:1]
	}
	if len(value) == 3 { // did not pick up the spare, e.g. 7/2
		firstTry := translate(value[0:1])
		secondTry := translate(value[2:3])
		return firstTry + secondTry
	}
	convertedNumber, _ := strconv.Atoi(value)
	return convertedNumber
}
