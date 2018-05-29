package main

import "fmt"

func FirstFactorial(num int64) int64 {
	if num == 0 {
		return 1
	}

	return num * FirstFactorial(num-1)
}

func main() {
	fmt.Println(FirstFactorial(4))
}
