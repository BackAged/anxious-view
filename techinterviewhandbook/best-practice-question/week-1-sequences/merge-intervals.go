// Problem: https://leetcode.com/problems/merge-intervals/

// Observation:
//            intersection only happens when ith start <= (i-1)th end and ith end > (i-1)th end

// Complexity:
//           time: O(N) - as we iterate the array only once
//           memory: 0(N) - N length array to store result

package main

import "sort"

// sorting ...
type SortIntervals [][]int

func (s SortIntervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortIntervals) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	} else if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}

	return false
}

func (s SortIntervals) Len() int {
	return len(s)
}

func merge(intervals [][]int) [][]int {
	// base case
	if len(intervals) <= 1 {
		return intervals
	}

	// sort the array
	s := SortIntervals(intervals)
	sort.Sort(s)

	ans := make([][]int, 0, len(intervals))

	start := intervals[0][0]
	end := intervals[0][1]

	index := 1

	for index < len(intervals) {
		startI := intervals[index][0]
		endI := intervals[index][1]

		switch {
		case startI <= end && endI > end:
			end = endI
		case startI > end:
			ans = append(ans, []int{start, end})

			start = startI
			end = endI

		}

		index++
	}

	ans = append(ans, []int{start, end})

	return ans

}
