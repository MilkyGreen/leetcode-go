package main

// 思路，先排序，从后面遍历，检查是否有符合的条件。排序每次取最大的值插到最后，这样如果有符合的能快速结束
// 思路完全错了，子数组意思是原数组中的一段，不能排序。。。
func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	maxIndex := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		temp := nums[maxIndex]
		nums[maxIndex] = nums[l-1-i]
		nums[l-1-i] = temp
		if nums[l-i-1] >= target {
			return 1
		}
	}

	tempValue := nums[l-1]
	tempLen := 1

	for i := l - 1; i > 0; i-- {
		if tempValue >= target {
			return tempLen
		}
		if nums[i] == nums[i-1]+1 {
			tempValue += nums[i-1]
			tempLen++
		} else {
			tempValue = 0
			tempLen = 0
		}
	}
	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

////**给定一个含有 n 个正整数的数组和一个正整数 target 。找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1,numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
// 滑动窗口：11   -->击败88% go提交
// [1,2,3,4,5]
func minSubArrayLen2(target int, nums []int) int {
	l := len(nums)
	left := 0
	right := 0
	ans := 0
	sum := nums[0]
	for left < l{
		if right == l-1 && sum<target{
			break
		}
		if sum >= target {
			if ans == 0 {
				ans = right-left + 1
			} else {
				ans = min(ans, right-left + 1)
			}
			sum -= nums[left]
			left++
		}else {
			right++
			sum += nums[right]
		}
	}
	return ans
}
