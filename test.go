package main

import (
	. "codewarrior/kata"
	"fmt"
)

func main() {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(Snail(snailMap))
	fmt.Println([]int{1, 2, 3, 6, 9, 8, 7, 4, 5})
}
