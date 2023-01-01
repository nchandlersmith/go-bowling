package bowling

import (
	"go-bowling/rules"
)

func Score(frames []string) int {
	score := 0
	strikeRule := rules.StrikeRule{}
	spareRule := rules.SpareRule{}
	countRule := rules.CountRule{}
	bowlingRules := []rules.Rule{strikeRule, spareRule, countRule}
	for i, turn := range frames {
		if isExtraTurn(i) {
			break
		}
		for _, rule := range bowlingRules {
			if rule.DoesApply(turn) {
				score += rule.Apply(i, frames)
				break
			}
		}
	}
	return score
}

func isExtraTurn(i int) bool {
	return i > 9
}
