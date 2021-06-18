package main

var area [][]byte
var rowNum int
var colNum int
// 遍历，遇到1时岛屿数量+1，dfs把所有相连的1改为2. 然后继续遍历
func numIslands(grid [][]byte ) int{
	ans := 0
	rowNum = len(grid)
	colNum = len(grid[0])
	area = grid
	for i:=0;i<rowNum;i++{
		for j:=0;j<colNum;j++{
			if area[i][j] == '1' {
				ans++
				infect(i,j)
			}
		}
	}
	return ans
}
// 感染所有相连的1，值改为2，这样后面遍历就会跳过
func infect(row int,col int){
	if row < 0 || row > rowNum-1 || col < 0 || col > colNum-1 {
		return
	}
	if area[row][col] == '1'{
		area[row][col] = '2'
		infect(row-1,col)
		infect(row+1,col)
		infect(row,col-1)
		infect(row,col+1)
	}
}