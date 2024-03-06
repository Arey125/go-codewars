package main

import (
	. "codewarrior/kata"
	"fmt"
	"sort"
)

func doTest(s string, expected []string) {
	actual := Permutations(s)
	sort.Strings(actual)

	fmt.Println()
	fmt.Println(actual)
	fmt.Println(expected)

	if len(actual) != len(expected) {
		fmt.Println("incorrect")
		return
	}

	for i, val := range actual {
		if expected[i] != val {
			fmt.Println("incorrect")
			return
		}
	}

	fmt.Println("correct")
}

func main() {
	doTest("a", []string{"a"})
	doTest("ab", []string{"ab", "ba"})
	doTest("abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"})
	abcd := []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
		"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
		"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
		"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}
	doTest("abcd", abcd)
	doTest("bcda", abcd)
	doTest("dcba", abcd)
	doTest("aa", []string{"aa"})
	doTest("aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"})
	doTest("aaaab", []string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"})
	doTest("aaaaab", []string{"aaaaab", "aaaaba", "aaabaa", "aabaaa", "abaaaa", "baaaaa"})
}
