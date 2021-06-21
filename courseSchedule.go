package main

// 思路：构建有向图，判断是否有环
func canFinish(numCourses int,prerequisites [][]int) bool{
	edge := make([][]int,numCourses) // 课程i -》依赖i的课程
	flags := make([]int,numCourses) // 课程状态 0未遍历 1遍历中 2已遍历
	for i:=0;i<numCourses;i++{
		edge[i] = make([]int,0)
	}
	for i:=0;i<len(prerequisites);i++{
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
	}
	for i:=0;i<numCourses;i++{
		if !dfs(edge,i,flags){ return false}
	}
	return true
}

func dfs(edges [][]int,i int,flags []int) bool{
	if flags[i] == 1{
		return false
	}
	if flags[i] == 2{
		return true
	}
	flags[i] = 1
	for _,j := range edges[i]{
		if !dfs(edges,j,flags){ return false}
	}
	flags[i] = 2
	return true
}