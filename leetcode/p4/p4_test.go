/*
	4.Median of Two Sorted Arrays

	There are two sorted arrays nums1 and nums2 of size m and n respectively.

	Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

	Example 1:
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0
	Example 2:
	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5

	**********************************************************

	4.两个排序数组的中位数

	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

	请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

	示例 1:

	nums1 = [1, 3]
	nums2 = [2]

	中位数是 2.0
	示例 2:

	nums1 = [1, 2]
	nums2 = [3, 4]

	中位数是 (2 + 3)/2 = 2.5
*/

package p4

import (
	"testing"
)

var cases = []struct {
	nums1  []int
	nums2  []int
	expect float64
}{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6, 8, 10}, 4},
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	f := FindMedianSortedArrays
	for _, c := range cases {
		if r := f(c.nums1, c.nums2); r != c.expect {
			b.Errorf(" \n%v\n%v\nfailed. Got %f, expected 3.0", c.nums1,c.nums2, r)
		}
	}
}
