package p2

import (
	"math"
	"testing"
)

var cases = []struct {
	l1, l2 *ListNode

	expect int
}{
	{
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		807,
	},
}

func calc(n *ListNode) (r int) {
	for i := 0; ; i++ {
		r = r + n.Val*int(math.Pow10(i))
		if n.Next != nil {
			n = n.Next
		} else {
			break
		}
	}
	return
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	f := AddTwoNumbers
	for _, c := range cases {
		if r := calc(f(c.l1, c.l2)); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}

func BenchmarkAddTwoNumbersFastest(b *testing.B) {
	f := AddTwoNumbersFastest
	for _, c := range cases {
		if r := calc(f(c.l1, c.l2)); r != c.expect {
			b.Errorf("result is %v, but expect is %v\n", r, c.expect)
		}
	}
}
