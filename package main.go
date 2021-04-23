package main

import (
	"strconv"
)

func main() {

}

func largestNumber(nums []int) string {
	// 排序，规则是：先转成字符串s1 s2 ，比较 s1+s2和s2+s1 大的排前面
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			s1 := strconv.Itoa(nums[j])
			s2 := strconv.Itoa(nums[j+1])
			if s1+s2 < s2+s1 {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
	s := ""
	if nums[0] == 0 {
		return "0"
	}
	for _, i := range nums {

		s += strconv.Itoa(i)
	}
	return s
}
