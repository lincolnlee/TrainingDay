package main

import (
	"fmt"
	"os"
	"strconv"
)

//工具类，负责提供一些通用的工具封装
type Utility struct {
}

func (this *Utility) checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(123)
	}
}

func (this *Utility) formatLength(l int64) (ret string) {
	lstr := strconv.FormatInt(l, 10)

	for i := len(lstr) + 1; i <= 20; i++ {
		ret += "0"
	}

	return ret + lstr
}
