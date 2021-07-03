package main

import (
	"fmt"
)

func productExceptSelf(nums []int)[]int{
	l := len(nums)
	leftP := make([]int,l)  //  左边乘积
	rightP := make([]int,l) // 右边乘积
	for i:=0;i<l;i++{
		if i == 0{
			leftP[0] = nums[0]
			rightP[l - 1] = nums[l-1]
		}else {
			leftP[i] = leftP[i-1] * nums[i]
			rightP[l - 1 - i] = rightP[l - i] * nums[l-1-i]
		}
	}
	for i:=0;i<l;i++{
		if i == 0{
			nums[0] = rightP[1]
		}else if i == l-1{
			nums[l-1] = leftP[l-2]
		}else {
			nums[i] = leftP[i-1] * rightP[i+1]
		}
	}
	return nums
}

func main(){
	nums := []int{1,2,3,4}
	nums = productExceptSelf(nums)
	fmt.Printf("%v",nums)
}