package main

func main()  {
	data := []int{1,1,2}
	println(removeDuplicates(data))
}

func removeDuplicates(nums []int) int {
	if len(nums)<=1{
		return len(nums)
	}

	res := 1
	i :=1
	for ; i<len(nums); i++{
		if nums[i] == nums[i-1]{
			continue
		}
		if res != i {
			nums[res] = nums[i]
		}
		res++
	}
	return res
}
