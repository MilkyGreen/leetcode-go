package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func main() {
	
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if denominator == 0 {
		return "NaN"
	}
	var buffer bytes.Buffer
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		buffer.WriteString("-")
	}

	num := int(math.Abs(float64(numerator)))
	denom := int(math.Abs(float64(denominator)))

	buffer.WriteString(strconv.Itoa(num / denom))
	num = num % denom
	if num == 0 {
		return buffer.String()
	}

	buffer.WriteString(".")

	m := make(map[int]int, 10)
	repeatPos := -1
	for {
		num *= 10
		pos, ok := m[num]
		if false == ok {
			m[num] = buffer.Len()
		} else {
			repeatPos = pos
			break
		}
		buffer.WriteString(strconv.Itoa(num / denom))
		//fmt.Println(buffer, len(buffer), num)
		num %= denom
		if num == 0 {
			break
		}
	}

	if repeatPos == -1 {
		return buffer.String()
	}
	res := buffer.String()
	return fmt.Sprintf("%s(%s)", res[0:repeatPos], res[repeatPos:])
}
