package bowling

import "strconv"

func Score(frames []string) int {
	var score int = 0
	for _, frame := range frames {
		if len(frame) == 2 {
			frame = frame[0:1]
		}
		frameScore, _ := strconv.Atoi(frame)
		score += frameScore
	}
	return score
}
