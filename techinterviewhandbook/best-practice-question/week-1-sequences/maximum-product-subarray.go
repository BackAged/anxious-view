// Problem: https://leetcode.com/problems/maximum-product-subarray/

// Observation:
//            contiguous non-empty subarray - that means we don't need n2 dp solution.
//            if at any point the product is zero, if we continue on the subarray, the product will always be zero, that means- we should start loooking for new subarray instead of continuing the subarray
//            for negative elements - we keep taking them in the subarray and update max, as even number of negative occurence would increase product while odd number will decrease product, to make sure we take the sub array with largest possiblity we iterate the array twice from 2 direction

// Example : [1, -2, 3, 4] => ([1] * -2 * 3 * 4) ([4 * 3] * -2 * 1)

// Complexity:
//           time: N + N => N (two iteration)
//           memory: N (answer array which could be max N)

package main

// maxInt returns maximum of the two numbers
func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProduct(nums []int) int {
	curProduct := 1
	maxProduct := nums[0]

	for _, v := range nums {
		curProduct *= v

		maxProduct = maxInt(curProduct, maxProduct)

		if curProduct == 0 {
			curProduct = 1

			continue
		}
	}

	curProduct = 1

	for index := len(nums) - 1; index >= 0; index-- {
		curProduct *= nums[index]

		maxProduct = maxInt(curProduct, maxProduct)

		if curProduct == 0 {
			curProduct = 1

			continue
		}

	}

	return maxProduct
}
