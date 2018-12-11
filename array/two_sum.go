//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import "fmt"

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.

//You may assume that each input would have exactly one solution, and you may not use the same element twice.

func main() {
	var a []int
	var b, x, y, z int = 0, 0, 0, 0
	fmt.Print("please input three digit:\n")
	fmt.Scanf("%d%d%d", &x, &y, &z)
	a = append(a, x, y, z)
	fmt.Print("please input the target:\n")
	fmt.Scanf("%d", &b)
	d := twoSum(a, b)
	if len(d) == 0 {
		fmt.Print("the data don't match\n")
	} else {
		fmt.Print("the data match the target\n")
		fmt.Println(d)
	}
}

func twoSum(nums []int, target int) []int {
	aa := []int{}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		bb, ok := m[nums[i]]
		if ok && bb != i {
			aa = []int{m[nums[i]], i}
			return aa
		}
		m[target-nums[i]] = i
	}
	return aa
}
