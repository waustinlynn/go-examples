package main

import (
	"fmt"
	"strconv"
)

func killKthBit(n int, k int) int {
	newInt := int64(n)
	bit := strconv.FormatInt(newInt, 2)
	newString := ""
	fmt.Println(string(bit[0]))
	pos := len(bit) - k
	for i:=0;i<len(bit);i++ {
		if(i == pos){
			newString = newString + "0"
			continue
		}
		newString = newString + string(bit[i])
	}
	//return int(strconv.Btoi64(newString, 2))
	fmt.Println(newString, pos)
	result, _ := strconv.ParseInt(newString, 2, 64)
	return int(result)
  }

func main() {
	fmt.Println(killKthBit(37, 3))
}