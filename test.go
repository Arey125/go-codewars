package main

import (
	. "codewarrior/kata"
	"fmt"
)

func main() {
		fmt.Println(ValidISBN10("1112223339"))
    fmt.Println(true)
    fmt.Println()
		fmt.Println(ValidISBN10("048665088X"))
    fmt.Println(true)
    fmt.Println()
		fmt.Println(ValidISBN10("1293000000"))
    fmt.Println(true)
    fmt.Println()
		fmt.Println(ValidISBN10("1234554321"))
    fmt.Println(true)
    fmt.Println()
		fmt.Println(ValidISBN10("1234512345"))
    fmt.Println(false)
    fmt.Println()
		fmt.Println(ValidISBN10("1293"))
    fmt.Println(false)
    fmt.Println()
		fmt.Println(ValidISBN10("X123456788"))
    fmt.Println(false)
    fmt.Println()
		fmt.Println(ValidISBN10("ABCDEFGHIJ"))
    fmt.Println(false)
    fmt.Println()
		fmt.Println(ValidISBN10("XXXXXXXXXX"))
    fmt.Println(false)
    fmt.Println()
		fmt.Println(ValidISBN10("000000000x"))
    fmt.Println(false)
    fmt.Println()
}
