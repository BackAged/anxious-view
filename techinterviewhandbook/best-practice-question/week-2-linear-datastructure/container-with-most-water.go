// Problem: https://leetcode.com/problems/container-with-most-water/

// Observation:
//            container size = (indexJ - indexi) * min(arr[j], arr[i])
//            when index are far, we have higher width, higher area
//            when min(arr[j], arr[i]) is higher we get higher area
//            optimal move is to start with widest x and move on the side which has shorter height

// Complexity:
//           time: N
//           memory: const

package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	xStart := 0
	xEnd := len(height) - 1

	maxContainer := -1

	for xStart < xEnd {
		yStart := height[xStart]
		yEnd := height[xEnd]

		maxContainer = max(maxContainer, (xEnd-xStart)*min(yStart, yEnd))

		if yStart < yEnd {
			xStart++
		} else {
			xEnd--
		}
	}

	return maxContainer
}
