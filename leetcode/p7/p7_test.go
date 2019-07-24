/*
	7.Reverse Integer

	Given a 32-bit signed integer, reverse digits of an integer.

	Example 1:

	Input: 123
	Output: 321
	Example 2:

	Input: -123
	Output: -321
	Example 3:

	Input: 120
	Output: 21
	Note:
	Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range:
	[−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

	**********************************************************

	7.反转整数

	给定一个 32 位有符号整数，将整数中的数字进行反转。

	示例 1:

	输入: 123
	输出: 321
	 示例 2:

	输入: -123
	输出: -321
	示例 3:

	输入: 120
	输出: 21
	注意:

	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/

package p7

import (
	"testing"
)

var cases = []struct {
	i      int
	expect int
}{
	{123, 321},
	{-123, -321},
}

func BenchmarkReverse(b *testing.B) {
	f := Reverse
	for _, c := range cases {
		if r := f(c.i); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}
