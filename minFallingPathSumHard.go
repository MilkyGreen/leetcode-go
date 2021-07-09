package main

func minFallingPathSum(arr [][]int) int {
	row := len(arr)
	col := len(arr[0])
	lastRowMinIndex := -1  // 同时找数组中第一小和第二小的方法
	secondMinIndex := -1
	for j:=0;j<col;j++{
		if lastRowMinIndex == -1 {
			lastRowMinIndex = j
		}else {
			if arr[0][j]< arr[0][lastRowMinIndex]{
				secondMinIndex = lastRowMinIndex
				lastRowMinIndex = j
			}else if secondMinIndex == -1{
				secondMinIndex = j
			}else if secondMinIndex != -1 && arr[0][j]< arr[0][secondMinIndex]{
				secondMinIndex = j
			}
		}
	}

	for i:= 1;i<row;i++{
		curMinIndex := -1
		curSecondMinIndex := -1
		for j:=0;j<col;j++{
			if j != lastRowMinIndex{
				arr[i][j] = arr[i][j] + arr[i-1][lastRowMinIndex]
			}else {
				arr[i][j] = arr[i][j] + arr[i-1][secondMinIndex]
			}

			if curMinIndex == -1 {
				curMinIndex = j
			}else {
				if arr[i][j]< arr[i][curMinIndex]{
					curSecondMinIndex = curMinIndex
					curMinIndex = j
				}else if curSecondMinIndex == -1{
					curSecondMinIndex = j
				}else if curSecondMinIndex != -1 && arr[i][j]< arr[i][curSecondMinIndex]{
					curSecondMinIndex = j
				}
			}
		}
		lastRowMinIndex = curMinIndex
		secondMinIndex = curSecondMinIndex
	}
	return arr[row -1][lastRowMinIndex]
}

func min(a int,b int)int{
	if a< b{
		return a
	}else {
		return b
	}
}

