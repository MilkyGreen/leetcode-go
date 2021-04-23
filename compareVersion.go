package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {

	var v1 = strings.Split(version1,".")
	var v2 = strings.Split(version2,".")

	l1 := len(v1)
	l2 := len(v2)

	var i int
	for i = 0;i<l1 && i<l2;i++{
		int1,_ := strconv.Atoi(v1[i])
		int2,_ := strconv.Atoi(v2[i])
		if int1 > int2 {
			return 1
		}else if int1 < int2{
			return -1
		}
	}
	
	if(l1 > l2){
		for ;i<l1;i++{
			int1,_ := strconv.Atoi(v1[i])
			if(int1 > 0){
				return 1;
			}
		}
	}else {
		for ;i<l2;i++{
			int2,_ := strconv.Atoi(v2[i])
			if(int2 > 0){
				return -1;
			}
		}
	}
	return 0
}

func main(){
	i := compareVersion("1.01","1.001")
	fmt.Printf("%v\n",i)
}