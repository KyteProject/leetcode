/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

package leetcode

// Brute Force O(n^²)
func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

// Hashmap O(n)
func twoSum2(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, j := range nums {
		comp := target - j
		if res, ok := hm[comp]; ok {
			return []int{res, i}
		}
		hm[j] = i
	}
	return []int{}
}
