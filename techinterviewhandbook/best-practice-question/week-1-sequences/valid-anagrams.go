// Problem: https://leetcode.com/problems/valid-anagram

// Observations:
//             Two strings are anagram of each other if they have same character frequency.

package main

func isAnagram(s string, t string) bool {
	// base check
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}

	for _, r := range s {
		if _, ok := m[r]; ok {
			m[r] += 1
		} else {
			m[r] = 1
		}
	}

	for _, r := range t {
		if v, ok := m[r]; !ok || v <= 0 {
			return false
		} else {
			m[r] -= 1
		}
	}

	return true
}
