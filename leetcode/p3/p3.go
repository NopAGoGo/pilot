/*
	3.Longest Substring Without Repeating Characters

	Given a string, find the length of the longest substring without repeating characters.

	Examples:

	Given "abcabcbb", the answer is "abc", which the length is 3.

	Given "bbbbb", the answer is "b", with the length of 1.

	Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

	**********************************************************

	3.无重复字符的最长子串

	给定一个字符串，找出不含有重复字符的最长子串的长度。

	示例：

	给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

	给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

	给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
*/

package p3

/*
	优化后的
	滑动窗口
	时间复杂度O(n):2n
	空间复杂度O(min(m,n))
*/
func LengthOfLongestSubstringOptimized(s string) int {
	l := len(s)
	rl := 0
	mp := map[uint8]int{}
	for i, j := 0, 0; j < l; j++ {
		q := s[j]
		if curr, e := mp[q]; e && curr > i {
			i = curr
		}
		tl := j - i + 1
		if tl > rl {
			rl = tl
		}
		mp[q] = j + 1
	}
	return rl
}

/*
	滑动窗口
	时间复杂度O(n):2n
	空间复杂度O(min(m,n))
*/
func LengthOfLongestSubstring(s string) int {
	l := len(s)
	rl, lt, i, j := 0, 0, 0, 0
	var q rune
	mp := map[rune]int{}
	for ; i < l && j < l; i++ {
		for ; j < l; j++ {
			q = rune(s[j])
			if curr, e := mp[q]; e && curr >= i {
				lt = j - curr
				if lt > rl {
					rl = lt
				}
				mp[q] = j
				i = curr
				j++
				break
			}
			mp[q] = j
			lt++
			if lt > rl {
				rl = lt
			}
		}
	}
	return rl
}

/*
	暴力法
	时间复杂度O(n3)
	空间复杂度O(min(m,n))
*/
func LengthOfLongestSubstringForce(s string) int {
	l := len(s)
	rLen := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if allUnique(s[i:j]) {
				rLenTmp := j - i
				if rLenTmp > rLen {
					rLen = rLenTmp
				}
			}
		}
	}
	return rLen
}

func allUnique(ss string) bool {
	mp := map[rune]int{}
	for i, r := range ss {
		if _, e := mp[r]; e {
			return false
		}
		mp[r] = i
	}
	return true
}

/*
	最快
	时间复杂度O(n3):n(n-1)2
	空间复杂度O(1)
*/
func LengthOfLongestSubstringFastest(s string) int {
	last, max, count := 1, 0, 0
	m := [128]int{}
	for index, value := range s {
		if m[value] < last {
			count += 1
		} else {
			last = m[value]
			if count > max {
				max = count
			}
			count = index - last + 1
		}
		m[value] = index + 1
	}
	if count > max {
		max = count
	}
	return max
}

