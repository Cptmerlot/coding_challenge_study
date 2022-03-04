package main

import "fmt"

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

	tcs := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{11, 7, 2, 15}, 9, []int{1, 2}},
		{[]int{11, 7, 14, 3, 5, 2, 15}, 8, []int{3, 4}},
	}

	for _, tc := range tcs {
		o := twoSum(tc.nums, tc.target)
		if o[0] != tc.expected[0] && o[1] != tc.expected[1] {
			fmt.Println("Wrong")
		} else {
			fmt.Println("Correct")
		}
	}

}

func twoSum(nums []int, target int) []int {

	for fi, fnum := range nums {
		if fnum > target {
			continue
		}
		for i := fi + 1; i < len(nums); i++ {
			if fnum+nums[i] == target {
				return []int{fi, i}
			}
		}
	}

	return []int{}
}

/*
	for fi, fnum := range nums {
		new_nums := nums[fi+1:]
		for si, snum := range new_nums {
			if fnum+snum == target {
				return []int{fi, si + fi + 1}
			}
		}
	}
*/

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
