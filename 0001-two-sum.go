package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	num2 := make(map[int]int, len(nums))
	for k,v := range nums{
		needV := target - v
		if index,ok := num2[needV]; ok{
			return []int{index, k}
		}
		num2[v] = k
	}
	return nil //不能省略
}