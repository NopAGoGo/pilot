/*
	1.Two Sum

	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].

	**********************************************************

	1.两数之和

	给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

	你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
*/

package p1

import (
	"reflect"
	"testing"
)

var cases = []struct {
	nums   []int
	target int
	expect []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
}

func BenchmarkTwoSumForce(b *testing.B) {
	f := TwoSumForce
	for _, c := range cases {
		if r := f(c.nums, c.target); !reflect.DeepEqual(r, c.expect) {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkTwoSumHashTwice(b *testing.B) {
	f := TwoSumHashTwice
	for _, c := range cases {
		if r := f(c.nums, c.target); !reflect.DeepEqual(r, c.expect) {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkTwoSumHashOnce(b *testing.B) {
	f := TwoSumHashOnce
	for _, c := range cases {
		if r := f(c.nums, c.target); !reflect.DeepEqual(r, c.expect) {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkTwoSumFastest(b *testing.B) {
	f := TwoSumFastest
	for _, c := range cases {
		if r := f(c.nums, c.target); !reflect.DeepEqual(r, c.expect) {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}
