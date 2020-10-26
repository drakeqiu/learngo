package main

import (
	"fmt"
	"github.com/drakeqiu/learngo/container/nonrepeating"
)

func main() {
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr("abcabcbb"))
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr("bbbbb"))
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr("pwwkew"))
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr(""))
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr("b"))
	fmt.Println(
		nonrepeating.LengthOfNoRepeatingSubStr("abcdef"))
}
