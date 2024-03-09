package main

import (
	. "codewarrior/kata"
	"fmt"
)

func main() {
	dotest(7, 3, 4)
	dotest(11, 19, 10)
	dotest(40, 3, 28)
	dotest(14, 2, 13)
	dotest(100, 1, 100)
}

func dotest(n, k, e int) {
	fmt.Println(JosephusSurvivor(n, k))
	fmt.Println(e)
	fmt.Println()
}
