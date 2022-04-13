// Problem: https://leetcode.com/problems/maximum-average-subarray-i/

// Observation:
//            Contiguous Subarray

// N*K solution:
//             We iterate i -> 0 to (N-K) and calculate average of ith to i+kth

// N solution:
//           sliding window, we iterate the array once, we slide the window ahead

package main

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	curWindowStart := 0

	curWindowSum := 0

	maxAvg := math.MaxFloat64 * -1

	for index, element := range nums {
		curWindowSum += element

		if index-curWindowStart == k-1 {
			maxAvg = math.Max(maxAvg, float64(curWindowSum)/float64(k))

			curWindowSum -= nums[curWindowStart]

			curWindowStart++
		}
	}

	return maxAvg
}
