package main

import "fmt"

func main() {

	// 使用 for 循环打印到三角 乘法口诀表
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, (i * j))
		}
		fmt.Println()
	}
}
