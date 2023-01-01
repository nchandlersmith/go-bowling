package rules

type StrikeRule struct{}

func (sr StrikeRule) DoesApply(turn string) bool {
	return turn == "x"
}

func (sr StrikeRule) Apply(score int, i int, frames []string) int {
	nextThrow := frames[i+1][0:1]
	var throwAfterNext string
	if len(frames[i+1]) == 1 {
		throwAfterNext = frames[i+2][0:1]
	} else {
		throwAfterNext = frames[i+1][1:2]
	}
	scoreNextThrow := Translate(nextThrow)
	scoreThrowAfterNext := Translate(throwAfterNext)
	score += 10 + scoreNextThrow + scoreThrowAfterNext
	return score

}
