package main

import (
	"fmt"
)

func main() {
	fmt.Println(find(16))
}

/*
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
示例 1:
输入: 16
输出: true
示例 2:
输入: 5
输出: false
链接：https://leetcode-cn.com/problems/power-of-four
*/
// 方法1，一直除以4，直到有余数，如果num等于0还没有余数，返回true
func isPowerOfFour(num int) bool {

	// num&(num-1)检查是否是二的幂
	// 101010...10 用16进制表示就是0xaaaaaaaa
	return num > 0 && num&(num-1) == 0 && num&0xaaaaaaaa == 0
}

/*
十进制 	二进制
2 	10
4 	100 （1 在第 3 位）
8 	1000
16 	10000（1 在第 5 位）
32 	100000
64 	1000000（1 在第 7 位）
128 	10000000
256 	100000000（1 在第 9 位）
512 	1000000000
1024 	10000000000（1 在第 11 位）
找一下规律： 4 的幂次方的数的二进制表示 1 的位置都是在奇数位。
这个特殊的数有如下特点：
    足够大，但不能超过 32 位，即最大为 1111111111111111111111111111111（ 31 个 1）
    它的二进制表示中奇数位为 1 ，偶数位为 0
    符合这两个条件的二进制数是：
1010101010101010101010101010101
*/
func find(num int) bool {
	//return num > 0 && num&(num-1) == 0 && num&0x55555555 == num
	return num > 0 && num&(num-1) == 0 && num&0b1010101010101010101010101010101 == num
}
