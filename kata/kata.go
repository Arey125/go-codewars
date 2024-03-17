package kata


func Snail(snailMap [][]int) []int {
	n := len(snailMap)
	m := len(snailMap[0])

	result := make([]int, 0, n*m)
	used := make([][]bool, n)

	for i := range used {
		used[i] = make([]bool, m)
	}

	x := 0
	y := 0

	dx := 0
	dy := 1

	rotate := func() {
		dx, dy = dy, -dx
	}

	for i := 0; i < n*m; i++ {
		result = append(result, snailMap[x][y])
        used[x][y] = true

		nextX := x + dx
		nextY := y + dy

		if nextX < 0 || nextY < 0 {
			rotate()
			x = x + dx
			y = y + dy
			continue
		}

		if nextX >= n || nextY >= m {
            rotate()
			x = x + dx
			y = y + dy
			continue
		}

		if used[nextX][nextY] {
			rotate()
			x = x + dx
			y = y + dy
			continue
		}

        x = nextX
        y = nextY
	}

	return result
}
