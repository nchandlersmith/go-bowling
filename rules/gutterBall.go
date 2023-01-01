package rules

type GutterBall struct{}

func (gb GutterBall) DoesApply(turn string) bool {
	return len(turn) == 2 && turn[1] == '-'
}

func (gb GutterBall) Apply(index int, frames []string) int {
	return CountThrow(frames[index][0:1])
}
