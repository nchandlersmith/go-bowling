package bowling

import "strconv"

func Score(frames []string) int {
	var score int = 0
	for _, frame := range frames {
		frameScore, _ := strconv.Atoi(frame)
		score += frameScore
	}
	return score
}
