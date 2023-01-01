package rules

type NotSpare struct{}

func (ns NotSpare) DoesApply(turn string) bool {
	return 3 == len(turn)
}

func (ns NotSpare) Apply(index int, frames []string) int {
	return CountThrow(frames[index][0:1]) + CountThrow(frames[index][2:3])
}
