package main

import "fmt"

/*

Given an integer array nums and two integers k and t, return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

Example 1:
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1, t = 2
Output: true

Example 3:
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false

Example 4:
Input: nums = [1,5,1,9,5], k = 2, t = 3
Output: true


Constraints:

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 104
0 <= t <= 231 - 1

*/

func main() {
	var answer bool
	answer = containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0)
	if !answer {
		fmt.Println("No Good 1")
	}
	fmt.Printf("This was %v\n", answer)
	answer = containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2)
	if !answer {
		fmt.Println("No Good 2")
	}
	fmt.Printf("This was %v\n", answer)
	answer = containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	if answer {
		fmt.Println("No Good 3")
	}
	fmt.Printf("This was %v\n", answer)
	answer = containsNearbyAlmostDuplicate([]int{1, 5, 1, 9, 5, 0, 1, 9}, 2, 3)
	if !answer {
		fmt.Println("No Good 4")
	}
	fmt.Printf("This was %v\n", answer)
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	condition1 := false
	condition2 := false
	for i, num := range nums {

		if !condition1 {
			if i+k < len(nums) {
				if num == nums[i+k] {
					condition1 = true
				}
			}
		}

		if !condition2 {
			if i+t < len(nums) {
				if num == nums[i+t] {
					condition2 = true
				}
			}
		}

		if condition1 && condition2 {
			return true
		}

	}

	return false
}

/*

	condition1 := false
	condition2 := false
	for fi, fnum := range nums {
		for si, snum := range nums {
			if fnum == snum {
				if (si - fi) == k {
					condition1 = true
				}
				if (si - fi) == t {
					condition2 = true
				}
			}
		}

		if condition1 && condition2 {
			return true
		}
	}

	return false

*/
