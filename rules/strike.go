package rules

type StrikeRule struct{}

func (sr StrikeRule) DoesApply(turn string) bool {
	return turn == "x"
}

func (sr StrikeRule) Apply(i int, frames []string) int {
	nextThrow := frames[i+1][0:1]
	var throwAfterNext string
	if len(frames[i+1]) == 1 {
		throwAfterNext = frames[i+2][0:1]
	} else {
		throwAfterNext = frames[i+1][1:2]
	}
	scoreNextThrow := CountThrow(nextThrow)
	scoreThrowAfterNext := CountThrow(throwAfterNext)
	return 10 + scoreNextThrow + scoreThrowAfterNext

}
