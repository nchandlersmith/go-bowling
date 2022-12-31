package bowling_test

import (
	"testing"

	"go-bowling"
)

func TestScore_ReturnsNumberOfKnockedDownPins(t *testing.T) {
	const expected int = 2
	line := []string{"1", "1"}
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_DashIsAMiss(t *testing.T) {
	const expected int = 3
	line := []string{"1-", "2-"}
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}
