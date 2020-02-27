package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 3}
	//nums[0], nums[1] = nums[1], nums[0]
	//fmt.Println(nums)

	res := permuteUnique(nums)
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	//res := make([][]int, 0)
	//sort.Ints(nums)
	//path := make([]int, 0)
	//used := make([]bool, len(nums))
	//backtrack(nums, path, len(nums), 0, used, &res)
	//return res

	res := make([][]int, 0)
	path := make([]int, 0)
	dataUsed := make(map[int]bool, 0)
	depthUsed := make([]bool, len(nums))
	backtrackNotUseSort(nums, path, len(nums), 0, dataUsed, depthUsed, &res)
	return res
}

func backtrack(nums, path []int, length, depth int, used []bool, res *[][]int) {
	if depth == length {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < length; i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		// 递归
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, path, length, depth+1, used, res)
		used[i] = false
		path = path[:len(path)-1]
	}
}

func backtrackNotUseSort(nums, path []int, length, depth int, dataUsed map[int]bool, depthUsed []bool, res *[][]int) {
	if depth == length {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < length; i++ {
		if depthUsed[i] {
			continue
		}
		du, exit := dataUsed[nums[i]]
		if exit && du == false {
			continue
		}
		path = append(path, nums[i])
		depthUsed[i] = true
		dataUsed[nums[i]] = true
		backtrackNotUseSort(nums, path, length, depth+1, dataUsed, depthUsed, res)
		depthUsed[i] = false
		dataUsed[nums[i]] = false
		path = path[:len(path)-1]
	}
}
