import "container/list"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (24.67%)
 * Total Accepted:    537.5K
 * Total Submissions: 2.2M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 *
 */
func reverse(x int) int {
	var rev int
	l := list.New()

	if x < 0 {
		for x != 0 {
			a := x % 10
			l.PushBack(a)
			x /= 10

		}
		for l.Len() != 0 {
			if rev > -214748364 {
				ll := l.Front()
				rev = rev*10 + ll.Value.(int)
				l.Remove(ll)

			} else if rev == -214748364 {
				if ll := l.Front(); ll.Value.(int) >= -7 {
					rev = rev*10 + ll.Value.(int)
					l.Remove(ll)

				} else {
					rev = 0
					break

				}

			} else {
				rev = 0
				break

			}

		}

	} else {
		for x != 0 {
			a := x % 10
			l.PushBack(a)
			x /= 10

		}
		for l.Len() != 0 {
			if rev < 214748364 {
				ll := l.Front()
				rev = rev*10 + ll.Value.(int)
				l.Remove(ll)

			} else if rev == 214748364 {
				if ll := l.Front(); ll.Value.(int) <= 7 {
					rev = rev*10 + ll.Value.(int)
					l.Remove(ll)

				} else {
					rev = 0
					break

				}

			} else {
				rev = 0
				break

			}

		}

	}
	return rev
}
