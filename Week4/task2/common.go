package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(123)
	}
}

func formatLength(l int64) (ret string) {
	lstr := strconv.FormatInt(l, 10)

	for i := len(lstr) + 1; i <= 20; i++ {
		ret += "0"
	}

	return ret + lstr
}
