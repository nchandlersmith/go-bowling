package bowling

import (
	"go-bowling/rules"
)

func Score(frames []string) int {
	score := 0
	strikeRule := rules.StrikeRule{}
	spareRule := rules.SpareRule{}
	for i, turn := range frames {
		switch true {
		case isExtraTurn(i):
			// no -op
		case strikeRule.DoesApply(turn):
			score = strikeRule.Apply(score, i, frames)
		case spareRule.DoesApply(turn):
			score = spareRule.Apply(score, i, frames)
		default:
			score += rules.Translate(turn)
		}
	}
	return score
}

func isExtraTurn(i int) bool {
	return i > 9
}
