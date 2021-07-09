package main

func minFallingPathSum(matrix [][]int) int {

	row := len(matrix)
	col := len(matrix[0])
	ans := 1000000
	if row == 1{
		for _,v := range matrix[0]{
			ans = min(ans,v)
		}
		return ans
	}
	
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
				if col == 1{
					matrix[i][j] = matrix[i-1][j] + matrix[i][j]
				}else {
					if j == 0{
						matrix[i][j] = min(matrix[i-1][j] ,matrix[i-1][j+1])+matrix[i][j]
					}else if j == col-1{
						matrix[i][j] = min(matrix[i-1][j] ,matrix[i-1][j-1])+matrix[i][j]
					}else{
						matrix[i][j] = min(min(matrix[i-1][j] ,matrix[i-1][j-1]),matrix[i-1][j+1])+matrix[i][j]
					}
				}			
			if i == row -1{
				ans = min(ans,matrix[i][j])
			}
		}
	}
	return ans
}

func min(a int,b int) int {
	if a < b{
		return a
	}else {
		return b
	}
}