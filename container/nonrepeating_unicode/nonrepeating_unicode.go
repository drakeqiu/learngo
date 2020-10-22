package main

import "fmt"

func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNoRepeatingSubStr("测试测试"))
	fmt.Println(
		lengthOfNoRepeatingSubStr("我的测试"))
	fmt.Println(
		lengthOfNoRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNoRepeatingSubStr("黑化肥挥发发灰会花飞，灰化肥挥发发黑会飞花"))
}
