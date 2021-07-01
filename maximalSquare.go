package main

// 求1组成的最大正方形,遍历法
func maximalSquare2(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])

	ans:=0
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if matrix[i][j] == '1'{
				ans = max(getSquare(matrix,i,j),ans)
			}
		}
	}
	return ans
}

func getSquare(matrix [][]byte,i int,j int) int{
	row := len(matrix)
	col := len(matrix[0])
	area := 1
	l := 1;// 长宽
	outter : for i+l < row && j+l < col && matrix[i+l][j] == '1' && matrix[i][j+l] == '1'{
		for w := j;w<=j+l;w++{
			if matrix[i+l][w] != '1'{
				break outter
			}
		}
		for h := i;h<=i+l;h++{
			if matrix[h][j+l] != '1'{
				break outter
			}
		}
		l++
		area = max(area,l*l)
	}
	return area
}

func max(a int,b int) int  {
	if a>b{
		return a
	}else {
		return b
	}
}

func min(a int,b int) int  {
	if a>b{
		return b
	}else {
		return a
	}
}

// 动态规划：dp[i][j]代表 以i，j为右下角的正方形最大面积，
func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int,row)
	ans:=0
	for i:=0;i<row;i++{
		dp[i] = make([]int,col)
		for j:=0;j<col;j++{
			if matrix[i][j] == '1'{
				if i != 0 && j != 0{
					dp[i][j] = min(dp[i-1][j],min(dp[i][j-1],dp[i-1][j-1]))+1  // 等于其上 左 左上 中最小值加一
				}else {
					dp[i][j] = 1
				}
				ans = max(dp[i][j] * dp[i][j],ans)
			}
		}
	}
	return ans
}