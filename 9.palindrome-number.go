import "strconv"

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (39.98%)
 * Total Accepted:    446.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false

	}
	if x/10 == 0 {
		return true

	}
	strInt := strconv.Itoa(x)

	//strInt := fmt.Sprintf("%d", x)

	for i := 0; i < len(strInt)/2; i++ {
		if strInt[i] != strInt[len(strInt)-i-1] {
			return false

		}

	}
	return true
}
