package rules

type CountRule struct{}

func (cr CountRule) DoesApply(turn string) bool {
	return len(turn) == 1 || len(turn) == 3 || len(turn) == 2 && turn[1] == '-'
}

func (cr CountRule) Apply(index int, frames []string) int {
	return Translate(frames[index])
}
