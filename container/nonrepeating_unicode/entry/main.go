package main

import (
	"fmt"
	"github.com/drakeqiu/learngo/container/nonrepeating_unicode"
)

func main() {
	fmt.Println(
		nonrepeating_unicode.LengthOfNoRepeatingSubStr("测试测试"))
	fmt.Println(
		nonrepeating_unicode.LengthOfNoRepeatingSubStr("我的测试"))
	fmt.Println(
		nonrepeating_unicode.LengthOfNoRepeatingSubStr("一二三二一"))
	fmt.Println(
		nonrepeating_unicode.LengthOfNoRepeatingSubStr("黑化肥挥发发灰会花飞，灰化肥挥发发黑会飞花"))
}
