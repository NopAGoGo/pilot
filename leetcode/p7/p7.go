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

func Reverse(x int) int {
	ret := 0
	for {
		v := x % 10
		x = x / 10
		if v == 0 && x == 0 {
			break
		}
		if ret > (1<<31-1)/10 || (ret == (1<<31-1)/10 && v > 7) {
			return 0
		}
		if ret < (-1<<31)/10 || (ret == (-1<<31)/10 && v < -8) {
			return 0
		}
		ret = ret*10 + v
	}
	return ret
}
