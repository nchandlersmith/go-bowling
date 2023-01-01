package bowling

import "strconv"

func Score(frames []string) int {
	score := 0
	for i, turn := range frames {
		if turn == "x" {
			scoreNextThrow := translate(frames[i+1][0:1])
			scoreThrowAfterNext := translate(frames[i+1][1:2])
			score += 10 + scoreNextThrow + scoreThrowAfterNext
			continue
		}
		if turn[1:2] == "/" && len(turn) == 2 {
			scoreNextThrow := translate(frames[i+1][0:1])
			score += 10 + scoreNextThrow
			continue
		}
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
