package kata

import "fmt"

func Solution(list []int) string {
	res := make([]byte, 0)

	for ind, val := range list {
		if ind == 0 {
			res = append(res, []byte(fmt.Sprint(val))...)
			continue
		}

		if ind == len(list)-1 {
			if res[len(res)-1] != '-' {
				res = append(res, ',')
			}

			res = append(res, []byte(fmt.Sprint(val))...)
			continue
		}

		if val-1 == list[ind-1] && val+1 == list[ind+1] {
			if res[len(res)-1] != '-' {
				res = append(res, '-')
			}
		}

		if val-1 != list[ind-1] {
			res = append(res, byte(','))
			res = append(res, []byte(fmt.Sprint(val))...)
			continue
		}

		if val+1 != list[ind+1] {
			if res[len(res)-1] != '-' {
				res = append(res, ',')
			}
			res = append(res, []byte(fmt.Sprint(val))...)
			continue
		}
	}

	return string(res)
}
