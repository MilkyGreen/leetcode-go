package main

// area = lenth * height
func maxArea(height []int) int{

	ans := 0
	i:=0
	l := len(height)
	j := l-1
	for i<j{
		h := min(height[i],height[j])
		lenth := j-i
		ans = max(ans,h*lenth)
		if height[i] > height[j] {
			j--
		}else {
			i++
		}
	}
	return ans
}
func max(a int,b int) int{
	if a > b{
		return a
	}else {
		return b
	}
}

func min(a int,b int) int{
	if a > b{
		return b
	}else {
		return a
	}
}