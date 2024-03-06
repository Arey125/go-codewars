package kata

import (
	"fmt"
	"math"
	"sort"
)

/*

not everyone can be boarded

*/

type Person struct {
	From int
	To   int
}

type PersonRecord struct {
	Index int
	Item  Person
}

func Order(level int, queue []Person) []int {
	res := make([]int, 0)

	sortedQueue := make([]PersonRecord, 0, len(queue))
	for ind, person := range queue {
		sortedQueue = append(sortedQueue, PersonRecord{Index: ind, Item: person})
	}

	sort.Slice(sortedQueue, func(i, j int) bool {
		return sortedQueue[i].Item.From < sortedQueue[j].Item.From
	})

	fmt.Println(sortedQueue)

	aboard := make([]int, 0)

	currentPosition := level
	targetPosition := level

	board := func(ind int) {
		currentPosition = sortedQueue[ind].Item.From
		if len(res) != 0 && res[len(res)-1] != currentPosition {
			res = append(res, currentPosition)
		}
		sortedQueue = removeFromSortedQueue(sortedQueue, ind)
		aboard = append(aboard, sortedQueue[ind].Item.To)
	}

	drop := func(ind int) {
		currentPosition = aboard[ind]
		if len(res) != 0 && res[len(res)-1] != currentPosition {
			res = append(res, currentPosition)
		}
		aboard = remove(aboard, ind)
	}

	for len(sortedQueue) != 0 {
		if len(aboard) == 0 {
			indexOfFirstInQueue := findFirstInQueue(sortedQueue)
			targetPosition = sortedQueue[indexOfFirstInQueue].Item.From
		}

		boardInd := findFirstToBoard(sortedQueue, currentPosition, targetPosition)
		if len(aboard) == 0 {
			board(boardInd)
			continue
		}

		dropInd := findFirstToDrop(aboard, currentPosition, targetPosition)

		distanceToBoarding := math.Abs(float64(sortedQueue[boardInd].Item.From - currentPosition))
		distanceToDropping := math.Abs(float64(aboard[dropInd] - currentPosition))

    // TODO
		if distanceToBoarding < distanceToDropping {
			board(boardInd)
		} else {
			drop(dropInd)
		}
	}

	return res
}

func removeFromSortedQueue(sortedQueue []PersonRecord, ind int) []PersonRecord {
	return append(sortedQueue[:ind], sortedQueue[ind+1:]...)
}

func remove(aboard []int, ind int) []int {
	return append(aboard[:ind], aboard[ind+1:]...)
}

func findFirstInQueue(sortedQueue []PersonRecord) int {
	res := 0
	for ind, val := range sortedQueue {
		if sortedQueue[res].Index > val.Index {
			res = ind
		}
	}

	return res
}

func findFirstToBoard(sortedQueue []PersonRecord, currentPosition, targetPosition int) int {
	fmt.Println("first to board")
	fmt.Println(sortedQueue)

	ind, found := sort.Find(len(sortedQueue), func(i int) int {
		return currentPosition - sortedQueue[i].Item.From
	})

	if targetPosition < currentPosition || found {
		fmt.Println("down", ind)
		return ind
	}

	fmt.Println("up", ind + 1)
	return ind + 1
}

func findFirstToDrop(aboard []int, currentPosition, targetPosition int) int {
	ind, found := sort.Find(len(aboard), func(i int) int {
		return aboard[i] - currentPosition
	})
	if targetPosition < currentPosition || found {
		return ind
	}

	return ind + 1
}
