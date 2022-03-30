// Problem: https://leetcode.com/problems/product-of-array-except-self/

// Observation:
//            If there's two zero value, all product will be zero
//            If there's one zero value on ith index, the ith index will have non zero product, other's will have 0

// Constraint:
//           Can't use division operator
//           Can use multiplication, addition, substraction operator

// N*N Solution:
//             Iterate the array
//             for ith index -> find prefix product of element 0 to i-1 and multiply suffix product of element (i+1, n)

// N solution:
//           We can have extra one array, that will hold the suffix product (product of elements from right)
//           The input array could hold the prefix product (product of elements from the right)

package main

func getSuffixProduct(curIndex int, suffix []int) int {
	if curIndex == len(suffix)-1 {
		return 1
	}

	return suffix[curIndex+1]
}

func getPrefixProduct(curIndex int, prefix []int) int {
	if curIndex == 0 {
		return 1
	}

	return prefix[curIndex-1]
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums), len(nums))

	suffix := make([]int, len(nums), len(nums))

	// suffix product
	for index := len(nums) - 1; index >= 0; index-- {
		if index == len(nums)-1 {
			suffix[index] = nums[index]

			continue
		}

		suffix[index] = nums[index] * suffix[index+1]
	}

	// prefix product
	for index := 0; index < len(nums); index++ {
		if index == 0 {
			nums[index] = nums[index]

			continue
		}

		nums[index] *= nums[index-1]
	}

	// fill up answer
	for index := 0; index < len(nums); index++ {
		ans[index] = getPrefixProduct(index, nums) * getSuffixProduct(index, suffix)
	}

	return ans
}
