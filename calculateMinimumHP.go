package main
func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
    if row == 0{
        return 0
    }
	col := len(dungeon[0])
    if col == 0{
        return 0
    }

	dp := make([][]int,row)
	for i:= 0;i<row;i++{
		dp[i] = make([]int, col)
	}
	// 从后向前算，最后一个位置如果大于0，则只需要一点血量，否则需要消耗血量+1
	if dungeon[row-1][col-1] > 0{
		dp[row-1][col-1] = 1
	}else {
		dp[row-1][col-1] = -dungeon[row-1][col-1]+1
	}

	// 最后一行
	for i:=col-2;i>=0;i--{

		if dungeon[row-1][i] >= dp[row-1][i+1]{
			dp[row-1][i] = 1
		}else {
			dp[row-1][i] = dp[row-1][i+1] - dungeon[row-1][i]
		}
	}

	// 最后一列
	for i:=row-2;i>=0;i--{
		if dungeon[i][col-1] >= dp[i+1][col-1]{
			dp[i][col-1]  = 1
		}else {
			dp[i][col-1]  =  dp[i+1][col-1] - dungeon[i][col-1]
		}
	}

	// 其余的格子,dp[i][j] 需要的血量，取决于dungeon[i][j]的值是否大于后面需要的最小血量，大于则需要1点血就够了。小于则要计算自己消耗的血量和后面需要的血量
	for i:=row-2;i>=0;i--{
		for j:=col-2;j>=0;j--{
			if dungeon[i][j] >= (min(dp[i+1][j],dp[i][j+1])){
				dp[i][j]   = 1
			}else {
				dp[i][j]  = (min(dp[i+1][j],dp[i][j+1])) - dungeon[i][j]
			}
		}
	}

	return dp[0][0]
}


func max(x int,y int) int{
	if x>=y {
		return x
	}
	return y
}

func min(x int,y int) int{
	if x>=y {
		return y
	}
	return x
}