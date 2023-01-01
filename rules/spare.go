package rules

type SpareRule struct{}

func (sr SpareRule) DoesApply(turn string) bool {
	return turn[1:2] == "/" && len(turn) == 2
}

func (sr SpareRule) Apply(score int, index int, frames []string) int {
	scoreNextThrow := Translate(frames[index+1][0:1])
	score += 10 + scoreNextThrow
	return score
}
