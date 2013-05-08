package main

import (
	"fmt"
	//"math"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strconv"
	"strings"
)

var dbConnString string = "root:ubuntu@tcp(127.0.0.1:3306)/deallog?charset=utf8"

func AddData(filename string, hour string) {
	fmt.Println("uv:", len(uv_00), "pv:", pv_00["pv"])
	fmt.Println("isp:", isp_00)
	fmt.Println("region:", region_00)
	fmt.Println("visitStatus:", visitStatus_00)

	db, _ := sql.Open("mysql", dbConnString)
	//checkErr(err)

	//插入数据
	stmt, _ := db.Prepare("INSERT analysislog SET domain=?,date=?,hour=?,uv=?,pv=?,isp=?,region=?,visitstatus=?")
	//checkErr(err)

	ispS := ""
	for key, value := range isp_00 {
		if ispS == "" {
			ispS = key + ":" + strconv.Itoa(value)
		} else {
			ispS += "," + key + ":" + strconv.Itoa(value)
		}
	}

	regionS := ""
	for key, value := range region_00 {
		if regionS == "" {
			regionS = key + ":" + strconv.Itoa(value)
		} else {
			regionS += "," + key + ":" + strconv.Itoa(value)
		}
	}

	visitStatusS := ""
	for key, value := range visitStatus_00 {
		if visitStatusS == "" {
			visitStatusS = key + ":" + strconv.Itoa(value)
		} else {
			visitStatusS += "," + key + ":" + strconv.Itoa(value)
		}
	}
	res, _ := stmt.Exec(strings.TrimRight(filename, ".log"), "2013-5-3", hour, len(uv_00), pv_00["pv"], ispS, regionS, visitStatusS)
	//checkErr(err)
	id, _ := res.LastInsertId()
	//checkErr(err)
	fmt.Println(id)

	db.Close()

	makeMap()

	runtime.GC()
}
