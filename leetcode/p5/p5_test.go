/*
	5.Longest Palindromic Substring

	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

	Example 1:

	Input: "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.
	Example 2:

	Input: "cbbd"
	Output: "bb"

	**********************************************************

	5.最长回文子串

	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

	示例 1：

	输入: "babad"
	输出: "bab"
	注意: "aba"也是一个有效答案。
	示例 2：

	输入: "cbbd"
	输出: "bb"
*/

package p5

import (
	"testing"
)

var cases = []struct {
	s  string
	expect string
}{
	{"sdfklacfsundayyadnuskleurihsd", "sundayyadnus"},
}

func BenchmarkLongestPalindromeForce(b *testing.B) {
	f:=LongestPalindromeForce
	for _, c := range cases {
		if r := f(c.s); r!= c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	f:=LongestPalindrome
	for _, c := range cases {
		if r := f(c.s); r!= c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLongestPalindromeManacher(b *testing.B) {
	f:=LongestPalindromeManacher
	for _, c := range cases {
		if r := f(c.s); r!= c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkLongestPalindromeFastest(b *testing.B) {
	f:=LongestPalindromeFastest
	for _, c := range cases {
		if r := f(c.s); r!= c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

