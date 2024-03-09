package kata

func JosephusSurvivor(n, k int) int {
    survivors := make([]int, n)
    for ind := range survivors {
        survivors[ind] = ind + 1
    }

    pos := 0
    for len(survivors) != 1 {
        pos = (pos + k - 1) % len(survivors)
        survivors = append(survivors[:pos], survivors[pos + 1:]...)
    }

    return survivors[0]
}
