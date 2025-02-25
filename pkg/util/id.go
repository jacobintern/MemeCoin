package util

import "strconv"

func IDParser(idStr string) (int64, error) {
	return strconv.ParseInt(idStr, 10, 64)
}
