// 使用递归函数从 10 打印到 1。

package main

import (
	"fmt"
)

func main() {
	printT(10)
	printrc(1)

}

func printT(n int) {
	num := 1
	if n >= num {
		fmt.Println(n)
		printT(n - 1)
	}

}

const (
	flag = 10
)

func printrc(n int) {
	if n > flag {
		return
	}
	fmt.Println(n)
	printrc(n + 1)
}
