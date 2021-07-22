package main

func change(amount int, coins []int) int {
	l := len(coins)
	dp := make([]int,amount+1)
	// 完全背包套路，先循环物品，再循环目标额
	dp[0] = 1
	for j:=0;j<l;j++{
		for i:=coins[j];i<=amount;i++{
			dp[i] += dp[i- coins[j]]
		}
	}
	return dp[amount]
}