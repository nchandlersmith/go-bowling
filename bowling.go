package bowling

import "strconv"

func Score(frames []string) int {
	var score int = 0
	for _, frame := range frames {
		if len(frame) == 2 {
			frame = frame[0:1]
		}
		if len(frame) == 3 {
			firstThrow, _ := strconv.Atoi(frame[0:1])
			secondThrow, _ := strconv.Atoi(frame[2:3])
			score += firstThrow + secondThrow
			continue
		}
		frameScore, _ := strconv.Atoi(frame)
		score += frameScore
	}
	return score
}
