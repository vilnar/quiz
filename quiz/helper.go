package main

import (
	"strconv"
)

func stringToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		// panic(err)
		return 0
	}
	return i
}
