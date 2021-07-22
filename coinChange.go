package main


func coinChange(coins []int, amount int) int {
	l := len(coins)
	dp := make([]int,amount+1)
	for i:=1;i<=amount;i++{
		dp[i] = amount+1 // 不一定每个金额都能凑出，先给一个无效值
	}

	dp[0] = 0 // 0元需要0个硬币就能满足
	for j:= 0;j< l;j++{
		for i:=1;i<=amount;i++{
			if coins[j]<=i{
				dp[i] = min(dp[i],dp[i-coins[j]]+1)
			}
		}
	}
	r := dp[amount]
	if r == amount+1{
		return -1
	}else {
		return r
	}
}

func min(a int,b int) int{
	if a<b{
		return a
	}else {
		return b
	}
}