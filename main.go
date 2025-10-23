package main

import (
	"fmt"
	"os"
)

func main() {
	var n, count int
	fmt.Println("введите число")
	fmt.Fscan(os.Stdin, &n)

	for n < 12307 {
		if n < 0 {
			n = n * (-1)
			count++
		}
		if n%7 == 0 {
			n = n * 39
			count++
		}
		if n%9 == 0 {
			n = n*13 + 1
			count++
			continue
		}
		if count == 0 {
			n = (n + 2) * 3
		}
	}
	fmt.Println(n)

}
