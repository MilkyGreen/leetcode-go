package main

func lengthOfLIS(nums []int) int {
    l:= len(nums)
    dp := make([]int,l)
    dp[0] = 1
	ans := 1
    for i:= 1;i<l;i++{
        dp[i] = 1
        for j:= i-1;j>=0;j--{
            if nums[i]>nums[j]{
                dp[i] = max(dp[j]+1,dp[i])
            }
        }
        ans = max(dp[i],ans)
    }
    return ans
}
func max(a int,b int)int {
    if a>b{
        return a
    }else{
        return b
    }
}

func main(){
	lengthOfLIS([]int{4,10,4,3,8,9})
}