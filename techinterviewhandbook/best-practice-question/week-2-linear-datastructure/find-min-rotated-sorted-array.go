// Problem: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// Observation:
//            Array is rotated between 1 and n time
//            1 2 3 => 3 1 2 => 2 3 1 => 1 2 3
//            1 2 3 4 5 => 5 1 2 3 4 => 4 5 1 2 3 => 3 4 5 1 2 => 2 3 4 5 1 => +1 2 3 4 5

//            n time rotation makes the array sorted again
//            rotated sorted array has two parts [increasing -> minimum element-> decreasing] ex-3 4 5 1 2
//            for any ith element if i+1th < ith then ans = i+1th
//            for any ith element if ith < i-1th then ans = ith
//            for any ith element how do we figure out which direction to go- increasing order or decreasing order to get minimum.
//            => for any ith element if ith > end that means we should for right to find the minimum
//               else we go left for minimum

// Complexity:
//           Time: log(N)
//           Memory: constant

package main

func findMin(nums []int) int {
	high := len(nums) - 1
	low := 0
	mid := 0

	for low <= high {
		mid = (high + low) >> 1

		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if mid >= 1 && nums[mid] < nums[mid-1] {
			return nums[mid]
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nums[mid]
}
