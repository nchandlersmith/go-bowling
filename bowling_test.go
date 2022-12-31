package bowling_test

import (
	"testing"

	"go-bowling"
)

func TestScore_ReturnsNumberOfKnockedDownPins(t *testing.T) {
	line := []string{"1/6", "7/1"}
	const expected int = 15
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_DashIsAMiss(t *testing.T) {
	line := []string{"1-", "2-"}
	const expected int = 3
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_ScoreSpare(t *testing.T) {
	line := []string{"8/", "7/1"}
	const expected int = 25
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}
