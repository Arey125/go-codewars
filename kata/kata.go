package kata

import (
	"strconv"
	"strings"
)

func ConvertIpToInt(addr string) int {
	res := 0

	for _, strVal := range strings.Split(addr, ".") {
		val, err := strconv.Atoi(strVal)

		if err != nil {
			panic(err.Error())
		}
		res = res*256 + val
	}

	return res
}

func IpsBetween(start, end string) int {
    startInt := ConvertIpToInt(start)
    endInt := ConvertIpToInt(end)

    return endInt - startInt
}
