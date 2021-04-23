package main

import (
	
	"fmt"
)
/**思路：出现新的0，意味着乘以10， 2*5=10 ，二的倍数肯定多于5，找出阶乘路径上有多少个5即可。
**/
func trailingZeroes(n int) int {
	if n < 5 {return 0}
	var count int
	for i:= 5;i<=n;i+=5{
		temp := i
		for temp%5 == 0{
			count++
			temp = temp/5
		}
	}
	return count
}

func main(){
	i := trailingZeroes(30)
	fmt.Printf("%v \n",i)
}