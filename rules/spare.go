package rules

type SpareRule struct{}

func (sr SpareRule) DoesApply(turn string) bool {
	return len(turn) == 2 && turn[1:2] == "/"
}

func (sr SpareRule) Apply(index int, frames []string) int {
	scoreNextThrow := CountThrow(frames[index+1][0:1])
	return 10 + scoreNextThrow
}
