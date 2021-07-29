package utils

import "strconv"

func StringToUint(src string) (uint, error) {
	srcInUint64, errConverting := strconv.ParseUint(src, 10, 64)
	if errConverting != nil {
		return 0, errConverting
	}
	return uint(srcInUint64), nil
}
