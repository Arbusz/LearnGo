/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (24.57%)
 * Total Accepted:    340.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	var mid float64
	if len1 != 0 && len2 != 0 {
		head1, tail1 := nums1[0], nums1[len1-1]
		head2, tail2 := nums2[0], nums2[len2-1]
		if head2 > tail1 || head1 > tail2 {
			if len1 > len2 {
				if (len1%2 == 0 && len2%2 == 0) || (len1%2 != 0 && len2%2 != 0) {
					if head2 > tail1 {
						mid = float64(nums1[(len1+len2)/2-1]+nums1[(len1+len2)/2]) / 2

					} else {
						mid = float64(nums1[(len1-len2)/2-1]+nums1[(len1-len2)/2]) / 2

					}

				} else {
					if head2 > tail1 {
						mid = float64(nums1[(len1+len2)/2])

					} else {
						mid = float64(nums1[(len1-len2)/2])

					}

				}

			}
			if len1 < len2 {
				if (len1%2 == 0 && len2%2 == 0) || (len1%2 != 0 && len2%2 != 0) {
					if head2 > tail1 {
						mid = float64(nums2[(len2-len1)/2-1]+nums2[(len2-len1)/2]) / 2

					} else {
						mid = float64(nums2[(len1+len2)/2-1]+nums2[(len1+len2)/2]) / 2

					}

				} else {
					if head2 > tail1 {
						mid = float64(nums2[(len2-len1)/2])

					} else {
						mid = float64(nums2[(len1+len2)/2])

					}

				}

			}
			if len1 == len2 && head2 > tail1 {
				mid = float64(head2+tail1) / 2

			}
			if len1 == len2 && head1 > tail2 {
				mid = float64(head1+tail2) / 2

			}

		} else {
			var num []int
			for i := 0; i < len1+len2; i++ {
				if len(nums1) != 0 && len(nums2) != 0 {
					if nums1[0] < nums2[0] {
						num = append(num, nums1[0])
						nums1 = nums1[1:]

					} else {
						num = append(num, nums2[0])
						nums2 = nums2[1:]

					}

				}
				if len(nums1) == 0 && len(nums2) != 0 {
					num = append(num, nums2[0])
					nums2 = nums2[1:]

				}
				if len(nums1) != 0 && len(nums2) == 0 {
					num = append(num, nums1[0])
					nums1 = nums1[1:]

				}

			}
			if (len1+len2)%2 != 0 {
				mid = float64(num[(len1+len2)/2])

			} else {
				mid = float64(num[(len1+len2)/2-1]+num[(len1+len2)/2]) / 2

			}

		}

	} else if len1 == 0 && len2 != 0 {
		if len2%2 == 0 {
			mid = float64(nums2[len2/2-1]+nums2[len2/2]) / 2

		} else {
			mid = float64(nums2[len2/2])

		}

	} else if len1 != 0 && len2 == 0 {
		if len1%2 == 0 {
			mid = float64(nums1[len1/2-1]+nums1[len1/2]) / 2

		} else {
			mid = float64(nums1[len1/2])

		}

	}
	return mid
}
