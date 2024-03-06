package main

import (
	. "codewarrior/kata"
	"fmt"
)

var testCounter int = 0

func tt(level int, queue []Person, want []int) {
	testCounter++
	fmt.Println(Order(level, queue))
	fmt.Println(want)
	fmt.Println()
}

func main() {
	puzzle := []Person{
		{From: 3, To: 2},
		{From: 5, To: 2},
		{From: 2, To: 1},
		{From: 2, To: 5},
		{From: 4, To: 3},
	}
	startingFloor := 1
	tt(startingFloor, puzzle, []int{2, 5, 4, 3, 2, 1})
	startingFloor = 1
	puzzle = []Person{{From: 5, To: 3}, {From: 4, To: 2}}
	solution := []int{5, 4, 3, 2}
	tt(startingFloor, puzzle, solution)

	startingFloor = 4
	puzzle = []Person{{From: 5, To: 3}, {From: 4, To: 2}}
	solution = []int{5, 4, 3, 2}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{{From: 4, To: 2}, {From: 5, To: 3}}
	solution = []int{4, 2, 5, 3}
	tt(startingFloor, puzzle, solution)

	startingFloor = 1
	puzzle = []Person{{From: 5, To: 1}, {From: 4, To: 3}, {From: 2, To: 1}}
	solution = []int{5, 4, 3, 2, 1}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{{From: 5, To: 1}, {From: 4, To: 3}, {From: 2, To: 1}, {From: 2, To: 4}, {From: 5, To: 2}}
	solution = []int{2, 4, 5, 4, 3, 2, 1}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{{From: 4, To: 3}, {From: 4, To: 3}, {From: 4, To: 3}, {From: 4, To: 3}, {From: 5, To: 1}, {From: 2, To: 3}, {From: 2, To: 4}, {From: 5, To: 2}}
	solution = []int{2, 3, 4, 3, 5, 2, 1}
	tt(startingFloor, puzzle, solution)

	tt(startingFloor, []Person{}, []int{})

	puzzle = []Person{
		{From: 3, To: 2}, // 3rd passenger
		{From: 5, To: 6}, // 4th passenger
		{From: 2, To: 1}, // 5th passenger
		{From: 2, To: 5}, // 1st passenger
		{From: 4, To: 3}, // 2nd passenger
	}
	solution = []int{2, 5, 4, 3, 2, 5, 6, 2, 1}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{
		{From: 3, To: 6}, // 16th passenger
		{From: 2, To: 6}, // 1st  passenger
		{From: 2, To: 6}, // 2nd  passenger
		{From: 2, To: 6}, // 3rd  passenger
		{From: 2, To: 6}, // 4th  passenger
		{From: 2, To: 6}, // 5th  passenger
		{From: 2, To: 6}, // 10th passenger
		{From: 2, To: 6}, // 11th passenger
		{From: 2, To: 6}, // 12th passenger
		{From: 2, To: 6}, // 13th passenger
		{From: 3, To: 2}, // 8th  passenger
		{From: 5, To: 6}, // 15th passenger
		{From: 2, To: 1}, // 9th  passenger
		{From: 2, To: 5}, // 14th passenger
		{From: 4, To: 3}, // 7th  passenger
		{From: 6, To: 1}, // 6th  passenger
	}
	solution = []int{2, 6, 4, 3, 2, 1, 2, 5, 6, 3, 6}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{
		{From: 3, To: 2}, // Al
		{From: 5, To: 2}, // Betty
		{From: 2, To: 1}, // Charles
		{From: 2, To: 5}, // Dan
		{From: 4, To: 3}, // Ed
	}
	solution = []int{2, 5, 4, 3, 2, 1}
	tt(startingFloor, puzzle, solution)

	puzzle = []Person{
		{From: 3900, To: 500},
		{From: 450000, To: 50000000000},
		{From: 1800, To: -400},
		{From: 4300, To: 500},
	}
	solution = []int{3900, 1800, 500, -400, 450000, 50000000000, 4300, 500}
	tt(2000, puzzle, solution)

	startingFloor = 5
	puzzle = []Person{{From: 5, To: 4}, // 1st passenger
		{From: 5, To: 3},  // 2nd passenger
		{From: 3, To: 4},  // 3rd passenger
		{From: 0, To: 2},  // 5th passenger
		{From: 3, To: -4}, // 4th passenger
		{From: 1, To: 2}}  // 6th passenger
	solution = []int{5, 4, 3, 4, 3, -4, 0, 1, 2}
	tt(startingFloor, puzzle, solution)
}
