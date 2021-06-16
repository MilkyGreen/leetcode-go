package main


/**
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

**/

// 解法1：一步一步的移动  --> 超出时间限制
func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		l := len(nums)
		nextNum := nums[l-1] // 要放置的值,初始放0，所以取最后一个数
		// 先取出nums[l-1],把nums[l-1] 放到 0 的位置，依次向后循环
		for j := 0; j < l; j++ {
			temp := nums[j]
			nums[j] = nextNum
			nextNum = temp
		}
	}
}

// 解法2 直接计算每个元素要放置的位置，比如移动3部，index+3就是新位置，如果超出数组长度，取余之后的数应该就是新位置 --> 通过，但是只击败6%的go提交
func rotate2(nums []int, k int) {
	num_location := make(map[int]int) // 元素、应该呆的位置
	l := len(nums)
	for i := 0; i < l; i++ {
		pos := i + k
		if pos > l-1 {
			pos = pos % l
		}
		num_location[pos] = nums[i]
	}

	for k, v := range num_location {
		nums[k] = v
	}
}

// 解法3 确定第一个元素未来的位置，将数组分成两份，再反过来拼接起来即可。 击败57% go提交
func rotate3(nums []int, k int) {
	if k != 0 {
		l := len(nums)
		if k%l == 0 {
			return
		}
		// 确定第一个元素应该呆的位置
		pos := 0 + k
		if pos > l-1 {
			pos = pos % l
		}
		if pos != l {
			slice1 := nums[:l-pos+1] // nums[start,end] 含头不含尾
			slice2 := nums[l-pos:]
			nums1 := append(slice2, slice1...)
			copy(nums, nums1)
		}
	}

}
