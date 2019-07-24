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

/*
	暴力法
	时间复杂度O(n2):n(n-1)
	空间复杂度O(1)
*/
func TwoSumForce(nums []int, target int) []int {
	var i, j int
	length := len(nums)
	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
	两遍哈希表
	时间复杂度O(n):2n
	空间复杂度O(n)
	一遍存一遍找，注意排除自己
*/
func TwoSumHashTwice(nums []int, target int) []int {
	mp := make(map[int]int)
	for index, v := range nums {
		mp[v] = index
	}
	for i, v := range nums {
		complement := target - v
		if j, e := mp[complement]; e && i != j {
			return []int{i, j}
		}
	}
	return []int{0, 0}
}

/*
	一遍哈希表
	时间复杂度O(n):n
	空间复杂度O(n)
	一边找一边存
*/
func TwoSumHashOnce(nums []int, target int) []int {
	mp := make(map[int]int)
	for j, v := range nums {
		complement := target - v
		if i, e := mp[complement]; e && i != j {
			return []int{i, j}
		}
		mp[v] = j
	}
	return []int{0, 0}
}

/*
	最快
	一遍哈希表
	时间复杂度O(n):n
	空间复杂度O(n)
	一边找一边存

	不用排除自己
*/
func TwoSumFastest(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		numToFind := target - v
		if _, e := m[numToFind]; e {
			return []int{m[numToFind], i}
		}
		m[v] = i
	}
	return nil
}
