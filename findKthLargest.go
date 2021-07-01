package main

// 思路：快排，
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums,0,len(nums)-1,k)
}

func quickSelect(nums []int,left int,right int, k int) int{
	sorted := partition(nums,left,right)
	if sorted == len(nums) - k{
		return nums[sorted]
	}else if sorted > len(nums) - k{
		return quickSelect(nums,left,sorted-1,k)
	}else {
		return quickSelect(nums,sorted+1,right,k)
	}
}

func partition(nums []int,left int,right int) int{
	n := nums[right] // 快排基数
	i := left
	for ;left < right;left++{
		if nums[left] <= n{
			swap(nums,left,i)
			i++
		}
	}
	swap(nums,right,i)
	return i;
}
func swap(nums []int,a int,b int){
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}