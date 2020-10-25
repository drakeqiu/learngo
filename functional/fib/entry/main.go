package main

import (
	"bufio"
	"fmt"
	"github.com/drakeqiu/learngo/functional/fib"
	"io"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	/*fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13
	fmt.Println(f()) // 21*/
	printFileContents(f)
}
