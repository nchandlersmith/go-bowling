package bowling_test

import (
	"testing"

	"go-bowling"
)

func TestScore(t *testing.T) {
	tests := []struct {
		name string
		line []string
		want int
	}{
		{
			name: "Returns number of knocked down pins",
			line: []string{"1/6", "7/1"},
			want: 15,
		},
		{
			name: "Dash is a miss",
			line: []string{"1-", "2-"},
			want: 3,
		},
		{
			name: "Score a spare",
			line: []string{"8/", "7/1"},
			want: 25,
		},
		{
			name: "Score a strike",
			line: []string{"x", "7/1"},
			want: 25,
		},
		{
			name: "Score a perfect game",
			line: []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
			want: 300,
		},
		{
			name: "10 turns of 9 and a miss",
			line: []string{"9-", "-9", "-9", "9-", "9-", "9-", "9-", "9-", "9-", "9-"},
			want: 90,
		},
		{
			name: "10 turns of 5 and spare plus 5",
			line: []string{"5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5"},
			want: 150,
		},
		{
			name: "strike then spare",
			line: []string{"x", "5/", "5/2"},
			want: 37,
		},
	}
	for _, test := range tests {
		got := bowling.Score(test.line)
		if got != test.want {
			t.Errorf("Test: %v. Got: %v Want: %v", test.name, got, test.want)
		}
	}
}
