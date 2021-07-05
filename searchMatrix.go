package main

func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix)
	col := len(matrix[0])

	i := row - 1
	j := 0

	for i >= 0 && j < col{
		if matrix[i][j] == target{  // compare left-bottom,if bigger,remove the bottomrow,if smaller remove the leftcol
			return true
		}else if matrix[i][j] > target{
			i--
		}else {
			j++
		}
	}
	return false
}