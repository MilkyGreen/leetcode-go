package main
import (
	"fmt"
)
func main() {
	ans := twoSum2([]int{2,7,11,15},9)
	fmt.Printf("%v\n",ans)
}


func twoSum(numbers []int, target int) []int {
	ans := make([]int,2,2)
	maps := make(map[int]int)
	for index, v := range numbers {
		maps[v] = index
	}

	for index, v := range numbers {
		leftNum,ok := maps[target - v]
		if ok {
			ans[0] = index +1
			ans[1] = leftNum +1
			break
		}
	}
	return ans
}

// 双指针方法
func twoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target{
			right--
		}else if sum < target{
			left++
		}else {
			return []int{left+1,right+1}
		}
	}
	return []int{-1,-1}
}