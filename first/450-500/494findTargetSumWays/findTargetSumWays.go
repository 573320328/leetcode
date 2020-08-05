package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 1, 1}
	res := findTargetSumWays(nums, 3)
	fmt.Println("res is ", res)
}

/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，
你都可以从 + 或 -中选择一个符号添加在前面。
返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
示例 1:
输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:
-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3
一共有5种方法让最终目标和为3。
注意:
数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。
*/

// fixme 这种方式是那里出错了啊
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	// 定义dp,dp表示选前i个元素,能达目标为j的组合数
	dp := make(map[int]map[int]int)

	// 初始化
	for i := 0; i <= length; i++ {
		dp[i] = make(map[int]int)
	}
	// 初值,当只有一个元素时
	for j := 0; j <= S; j++ {
		if j == 0 && nums[0] == 0 {
			dp[0][j] = 1
		} else {
			if nums[0] == j {
				dp[0][j] += 1
			} else if nums[0] == -j {
				dp[0][j] += 1
			}
		}
	}
	// 状态转移方程
	fmt.Println("first dp ", dp)

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= S; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i]] + 1 // 选加号
			}
		}
	}
	fmt.Println("second dp ", dp[length-1][S])

	for i := 1; i < len(nums); i++ {
		for j := S; j >= 0; j-- {
			if j+nums[i] <= S {
				dp[i][j] += dp[i-1][j+nums[i]] + 1 // 选减号
			}
		}
	}
	fmt.Println("third dp ", dp[length-1][S])
	return dp[length-1][S]
}

func findTargetSumWays0(nums []int, S int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	if sum < int(math.Abs(float64(S))) || (sum+S)%2 != 0 {
		return 0
	}

	target := (sum + S) / 2
	dp := make(map[int]int)
	dp[0] = 1

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[target]
}