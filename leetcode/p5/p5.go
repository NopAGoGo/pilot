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

/*
	最长公共子串 暴力法
	时间复杂度O(n3)
	空间复杂度O(n):2n
*/
func LongestPalindromeForce(s string) string {
	sr := ""
	for _, v := range s {
		sr = string(v) + sr
	}
	n := len(s)
	start, longest := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := i, j
			match := 0
			for x < n && y < n {
				if s[x] != sr[y] {
					break
				}
				match++
				x++
				y++
			}
			if match > longest && i+match == n-j {
				longest = match
				start = i
			}
		}
	}
	return s[start : start+longest]
}

/*
	动态规划-中心扩展算法
	时间复杂度O(n2)
	空间复杂度O(1)
*/
func LongestPalindrome(s string) string {
	// n2
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		length := 0
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		if len1 > len2 {
			length = len1
		} else {
			length = len2
		}
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

/*
	Manacher算法
	时间复杂度O(n)
	空间复杂度O(n):2n
*/
func LongestPalindromeManacher(s string) string {
	// 处理字符串，插值适应奇偶字符串，统一后续计算方法
	// 前加"^"后加"$"，避免溢出处理
	t := preProcess(s)
	// 取字符串长度
	n := len(t)
	// 创建p切片
	p := make([]int, n)
	// center为当前下标,right为右边界下标
	center, right := 0, 0
	// 跳过前后符号进行循环
	for i := 1; i < n-1; i++ {
		iMirror := 2*center - i
		if right > i {
			if right-i > p[iMirror] {
				p[i] = p[iMirror]
				continue
			} else {
				p[i] = right - i
			}
		} else {
			p[i] = 0
		}

		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	index := (centerIndex - 1 - maxLen) / 2
	return s[index : index+maxLen]
}

// Transform s into t.
// For example, s = "abba", t = "^#a#b#b#a#$".
// ^ and $ signs are sentinels appended to each end to avoid bounds checking
func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}
	ret := "^"
	for _, v := range s {
		ret += "#" + string(v)
	}
	ret += "#$"
	return ret
}

/*
	最快
	也是动态规划-中心扩展算法
	时间复杂度O()
	空间复杂度O()
*/
func LongestPalindromeFastest(s string) string {
	// copy fastest
	str := []byte(s)
	if len(s) < 2 {
		return s
	}
	index := 0
	maxLen, start := 1, 0
	for index < len(s) {
		if maxLen/2 >= len(s)-index {
			break
		}
		j, k := index, index
		for k < len(s)-1 && str[k+1] == str[k] {
			k++
		}
		index = k + 1
		for k < len(s)-1 && j > 0 && str[j-1] == str[k+1] {
			j--
			k++
		}
		if k-j+1 > maxLen {
			start = j
			maxLen = k - j + 1
		}
	}
	return string(str[start : start+maxLen])
}
