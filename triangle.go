package main
import "fmt"
func minimumTotal(triangle [][]int) int {
	l:=len(triangle)
	if l == 1{
		return triangle[0][0]
	}
	for i:=1;i<l;i++{
		m := len(triangle[i])
		for j :=0;j<m;j++{
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			}else if j == m-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			}else {
				triangle[i][j] = min(triangle[i-1][j],triangle[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	bottom := triangle[l-1]
	ans := bottom[0]
	for _,v := range bottom{
		ans = min(ans,v)
	}

	return ans
}

func min(a int,b int) int {
	if a<b{
		return a
	}else {
		return b
	}
}

func main(){
	nums := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	ans := minimumTotal(nums)
	fmt.Printf("%v",ans)
}