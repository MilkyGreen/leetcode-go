package main

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]
	for i:=1;i<len(prices);i++ {
		max = getmax(max,prices[i] - min)
		min = getmin(prices[i], min)
	}
    return max
}

func getmin(a int,b int) int {
	if a<b{
		return a
	}else{
		return b
	}
}

func getmax(a int,b int) int {
	if a<b{
		return b
	}else{
		return a
	}
}

// 最多交易两次
func maxProfit3(prices []int) int {
	
}

// 无限次交易
func maxProfit2(prices []int) int {

    profit := 0
    min := prices[0]
    for i:=1;i<len(prices);i++{
		if prices[i] > min{
			profit += prices[i]-min
			min = prices[i]
		}
		min = getmin(min,prices[i])
	}
	return profit

}

