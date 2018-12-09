/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (29.44%)
 * Total Accepted:    259.4K
 * Total Submissions: 881.2K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */
func convert(s string, numRows int) string {
	if len(s) <= numRows || len(s) == 1 || numRows <= 1 {
		return s

	}
	strArr := make([]byte, 0)
	stepLen := 2*numRows - 2
	//令p=numRows×2-2，可以总结出以下规律
	//0行， 0×p，1×p，...
	//r行， r，1×p-r，1×p+r，2×p-r，2×p+r，...
	//最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p
	//即对于第i行，如果i非首行和末行，除了s[kp+i]外，还需要s[(k+1)p-i]
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += stepLen {
			strArr = append(strArr, s[i+j])
			if i != 0 && i != numRows-1 && j+stepLen-i < len(s) {
				strArr = append(strArr, s[j+stepLen-i])

			}

		}

	}
	return string(strArr)
}
