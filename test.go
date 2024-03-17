package main

import (
	. "codewarrior/kata"
	"fmt"
)

func main() {
	s := Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
	fmt.Println(s)
    fmt.Println("-6,-3-1,3-5,7-11,14,15,17-20")
	fmt.Println()

    fmt.Println(Solution([]int{1, 2, 4, 5, 7, 8}))
	fmt.Println("1,2,4,5,7,8")
}
