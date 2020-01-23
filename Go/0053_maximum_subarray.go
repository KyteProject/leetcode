/*
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
*/

import "math"

// LinearTime Dynamic Programming (Kadane's algorithm) O(n^0)
func maxSubArray(nums []int) int {
	n := len(nums)
	maxSum := nums[0]
	for i := 1; i < n; i++ {
			if nums[i - 1] > 0 {
					nums[i] += nums[i - 1]
			}
			if maxSum < nums[i] {
					maxSum = nums[i]
	}
	}
	return maxSum
}

// Greedy Linear Time O(n)
func maxSubArray(nums []int) int {
	n := len(nums)
	currSum := nums[0]
	maxSum := nums[0]
	
	for i := 1; i < n; i++ {
			if nums[i] > currSum + nums[i] {
					currSum = nums[i]
			} else {
					currSum = currSum + nums[i]
			}
	 
			if maxSum < currSum {
					maxSum = currSum
			}
	}
	return maxSum
}

// Quadratic Time Bruteforce O(n^2)
func maxSubArray(nums []int) int {
	n := len(nums)
	maxSum := math.MinInt32
	
	for i := 0; i < n; i++ {
			currentSum := 0
			for j := i; j < n; j++ {
					currentSum += nums[j]
					
					if maxSum < currentSum {
							maxSum = currentSum
					}
			}
	}
	return maxSum
}

// Cubic Time Bruteforce O(n^3)
func maxSubArray(nums []int) int {
    n := len(nums)
    maxSum := math.MinInt32
    
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            sum := 0
            
            for k := i; k <= j; k++ {
                sum += nums[k]
            }
            
            if maxSum < sum {
                maxSum = sum
            }
        }
    }
    return maxSum
}