package bowling_test

import (
	"testing"

	"go-bowling"
)

func TestScore_ReturnsNumberOfKnockedDownPins(t *testing.T) {
	line := []string{"1/6", "7/1"}
	const expected = 15
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_DashIsAMiss(t *testing.T) {
	line := []string{"1-", "2-"}
	const expected = 3
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_ScoreSpare(t *testing.T) {
	line := []string{"8/", "7/1"}
	const expected = 25
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual: '%v", expected, actual)
	}
}

func TestScore_ScoreStrike(t *testing.T) {
	line := []string{"x", "7/1"}
	const expected = 25
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual '%v'", expected, actual)
	}
}

func TestScore_PerfectGame(t *testing.T) {
	line := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	const expected = 300
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual '%v'", expected, actual)
	}
}

func TestScore_10TurnsOf9AndMiss(t *testing.T) {
	line := []string{"9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-"}
	const expected = 90
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual '%v'", expected, actual)
	}
}

func TestScore_10TurnsOf5AndSparePlus5(t *testing.T) {
	line := []string{"5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5"}
	const expected = 150
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual '%v'", expected, actual)
	}
}

func TestScore_StrikeThenSpare(t *testing.T) {
	line := []string{"x", "5/", "5/2"}
	const expected = 37
	actual := bowling.Score(line)
	if actual != expected {
		t.Errorf("expected: '%v' actual '%v'", expected, actual)
	}
}
