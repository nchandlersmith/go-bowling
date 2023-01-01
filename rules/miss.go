package rules

type Miss struct{}

func (m Miss) DoesApply(turn string) bool {
	return len(turn) == 2 && (turn[1] == '-' || turn[0] == '-')
}

func (m Miss) Apply(index int, frames []string) int {
	return CountThrow(frames[index][0:1]) + CountThrow(frames[index][1:2])
}
