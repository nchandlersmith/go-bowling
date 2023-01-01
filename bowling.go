package bowling

import "strconv"

func Score(frames []string) int {
	score := 0
	for i, turn := range frames {
		switch true {
		case isExtraTurn(i):
			// no -op
		case isStrike(turn):
			score = scoreStrike(frames, i, score)
		case isSpare(turn):
			score = scoreSpare(frames, i, score)
		default:
			score += translate(turn)
		}
	}
	return score
}

func isExtraTurn(i int) bool {
	return i > 9
}

func isStrike(turn string) bool {
	return turn == "x"
}

func scoreStrike(frames []string, i int, score int) int {
	nextThrow := frames[i+1][0:1]
	var throwAfterNext string
	if len(frames[i+1]) == 1 {
		throwAfterNext = frames[i+2][0:1]
	} else {
		throwAfterNext = frames[i+1][1:2]
	}
	scoreNextThrow := translate(nextThrow)
	scoreThrowAfterNext := translate(throwAfterNext)
	score += 10 + scoreNextThrow + scoreThrowAfterNext
	return score
}

func isSpare(turn string) bool {
	return turn[1:2] == "/" && len(turn) == 2
}

func scoreSpare(frames []string, i int, score int) int {
	scoreNextThrow := translate(frames[i+1][0:1])
	score += 10 + scoreNextThrow
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
	if "x" == value {
		return 10
	}
	convertedNumber, _ := strconv.Atoi(value)
	return convertedNumber
}
