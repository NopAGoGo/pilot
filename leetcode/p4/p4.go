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

/*
	逻辑推理+二分搜索
	时间复杂度O(log2min(m,n))
	空间复杂度O(1)
*/
func FindMedianSortedArrays(a []int, b []int) float64 {
	m, n := len(a), len(b)
	if m > n {
		a, b = b, a
		m, n = n, m
	}
	maxLeft := 0
	minRight := 0
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && b[j-1] > a[i] {
			iMin = i + 1
		} else if i > iMin && a[i-1] > b[j] {
			iMax = i - 1
		} else {
			if i == 0 {
				maxLeft = b[j-1]
			} else if j == 0 {
				maxLeft = a[i-1]
			} else {
				if a[i-1] > b[j-1] {
					maxLeft = a[i-1]
				} else {
					maxLeft = b[j-1]
				}
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				minRight = b[j]
			} else if j == n {
				minRight = a[i]
			} else {
				if a[i] < b[j] {
					minRight = a[i]
				} else {
					minRight = b[j]
				}
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return -1.0
}
