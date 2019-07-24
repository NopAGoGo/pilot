package p3

import (
	"testing"
)

var cases = []struct {
	s      string
	expect int
}{
	{"abcabcbb", 3},
	{"uqinntq", 4},
	{"abba", 2},
	{"dvdf", 3},
	{"", 0},
	{"a", 1},
	{"au", 2},
	{"aab", 2},
	{"cdd", 2},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", 94},
}

func BenchmarkLengthOfLongestSubstringOptimized(b *testing.B) {
	f := LengthOfLongestSubstringOptimized
	for _, c := range cases {
		if r := f(c.s); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	f := LengthOfLongestSubstring
	for _, c := range cases {
		if r := f(c.s); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLengthOfLongestSubstringForce(b *testing.B) {
	f := LengthOfLongestSubstringForce
	for _, c := range cases {
		if r := f(c.s); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLengthOfLongestSubstringFastest(b *testing.B) {
	f := LengthOfLongestSubstringFastest
	for _, c := range cases {
		if r := f(c.s); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}
