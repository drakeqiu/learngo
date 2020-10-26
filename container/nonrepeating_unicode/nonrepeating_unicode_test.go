package nonrepeating_unicode

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

		// Chinese support
		{"测试测试", 2},
		{"这里是百度", 5},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := LengthOfNoRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkLengthOfNoRepeatingSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8
	b.ResetTimer()

	b.Logf("len(s) = %d", len(s))
	for i := 0; i < b.N; i++ {
		actual := LengthOfNoRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",
				actual, s, ans)
		}
	}
}
