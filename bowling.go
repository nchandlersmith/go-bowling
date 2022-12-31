package bowling

import "strconv"

func Score(frames []string) int {
	score := 0
	for _, frame := range frames {
		score += translate(frame)
	}
	return score
}

func translate(value string) int {
	if len(value) == 2 { // second throw was a miss, e.g. 7-
		value = value[0:1]
	}
	if len(value) == 3 { // did not pick up the spare, e.g. 7/2
		firstThrow := translate(value[0:1])
		secondThrow := translate(value[2:3])
		return firstThrow + secondThrow
	}
	convertedNumber, _ := strconv.Atoi(value)
	return convertedNumber
}
