package bowling

import "strconv"

func Score(frames []string) int {
	score := 0
	for _, frame := range frames {
		if len(frame) == 2 {
			frame = frame[0:1]
		}
		if len(frame) == 3 {
			firstThrow := notationToScore(frame[0:1])
			secondThrow := notationToScore(frame[2:3])
			score += firstThrow + secondThrow
			continue
		}
		score += notationToScore(frame)
	}
	return score
}

func notationToScore(value string) int {
	convertedNumber, _ := strconv.Atoi(value)
	return convertedNumber
}
