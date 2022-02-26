package main

import (
	"fmt"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]

input: nums =  [11,7,2,15], target = 9
Output: [1,2]

input: nums =  [11,7,14,3,5,2,15], target = 8
Output: [3,4]

*/

func main() {

	o := twoSum([]int{11, 7, 14, 3, 5, 2, 15}, 8)
	fmt.Printf("%+v", o)
}

func twoSum(nums []int, target int) []int {

	for fi, fnum := range nums {
		new_nums := nums[fi+1:]
		for si, snum := range new_nums {
			if fnum+snum == target {
				return []int{fi, si + fi + 1}
			}
		}
	}

	return []int{}
}

/*
	[11,7,2,15] = 9
	for fi, fnum := range nums {
		for si, snum := range nums {
			if fi == si {
				continue
			}
			if fnum+snum == target {
				return []int{fi, si}
			}
		}
	}

*/

/*

	nums = sort.IntSlice(nums)

	leftIndex := 0
	rightIndex := -1
	for {

		t := nums[leftIndex] + nums[len(nums)+rightIndex]
		switch {
		case t > target:
			rightIndex = rightIndex - 1
		case t < target:
			leftIndex = leftIndex + 1
		case t == target:
			return []int{leftIndex, len(nums) + rightIndex}
		}

	}

*/
