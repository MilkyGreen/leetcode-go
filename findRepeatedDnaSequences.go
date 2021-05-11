package main

import (
	"fmt"
)

//思路：宽度为10的滑动窗口，遍历字符串，将每个子串保存为map的key，次数为value，最后输出次数大于1的字符串列表
func findRepeatedDnaSequences(s string) []string {
	rs := []rune(s)
	l := len(rs)
	strs := make(map[string] int)
	for i:=0;i<=l-10;i++{
		substr := (string)(rs[i:i+10])
		_,ok :=strs[substr]
		if ok {
			strs[substr] = 1 + strs[substr]
		}else {
			strs[substr] = 1
		}
	}
	res := make([]string,0)
	for k,v := range strs{
		if v > 1{
			res = append(res,k)
		}
	}
	return res
}

func main(){
	s := findRepeatedDnaSequences("AAAAAAAAAAA")
	fmt.Println("V%",s)
}