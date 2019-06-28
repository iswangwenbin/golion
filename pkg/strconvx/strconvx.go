package strconvx

import (
	"strconv"
	"strings"
)

// SplitToUint split the string to uint64 by sep
func SplitToUint(str string, sep string) []uint64 {
	var result []uint64
	data := strings.Split(str, sep)
	for _, v := range data {
		s, _ := strconv.ParseUint(v, 10, 64)
		result = append(result, s)
	}
	return result
}
