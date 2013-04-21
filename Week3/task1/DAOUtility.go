package main

import (
	"fmt"
	"os"
	"time"
)

var dbConnString string = "root:ubuntu@tcp(127.0.0.1:3306)/deallog?charset=utf8"

var monthChange = map[string]int{
	"Jan": 1,
	"Feb": 2,
	"Mar": 3,
	"Apr": 4,
	"May": 5,
	"Jun": 6,
	"Jul": 7,
	"Aug": 8,
	"Sep": 9,
	"Oct": 10,
	"Nov": 11,
	"Dec": 12,
}

type RssSource struct {
	Id      int
	RssName string
	RssUrl  string
}

type RssContent struct {
	Id      int
	Atitle  string
	Adesc   string
	Alink   string
	Atime   time.Time
	Srcid   int32
	Ahash1  uint32
	Ahash2  uint32
	SrcName string
}

type PageInfo struct {
	IsCurrent bool
	PageIdx   int
}

type RssContentPage struct {
	PageIdx      int
	Page         *[]PageInfo
	RssContentLs *[]RssContent
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
