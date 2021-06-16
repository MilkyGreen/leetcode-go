package main


//[2,7,9,3,1] 输出：12 。 动态规划：房间i的最大值=max(自己+i-2房间最大值，i-1房间最大值)，可以倒着思考，选i就不能选i-1，因此是i加上i-2的最大值。

func rob(nums []int) int {
	l := len(nums)
	dp := make([]int,l) // 到每个房间的最大值，不一定选这个房间
	dp[0] = nums[0]
	if l == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0],nums[1]) // 注意这一行必须在if外面...
	if l == 2{
		return dp[1]
	}
	for i:=2;i<l;i++{
		dp[i] = max(nums[i] + dp[i-2],dp[i-1])
	}
	return dp[l - 1]
}

func max(i int,j int) int{
	if i > j {
		return i
	}else {
		return j
	}
}

// 优化版，只需要保存前两个房间的最大值,每走一步更新一下两个值
func rob2(nums []int) int {
	l := len(nums)
	first := nums[0]
	if l == 1 {
		return first
	}
	secend := max(nums[0],nums[1]) 
	if l == 2{
		return secend
	}
	for i:=2;i<l;i++{
		first,secend = secend,max(nums[i] + first,secend)
	}
	return secend
}