package nonrepeating

import "testing"

func TestLengthOfNoRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcabcd", 4},
	}

	for _, tt := range tests {
		actual := LengthOfNoRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.ans)
		}
	}
}
