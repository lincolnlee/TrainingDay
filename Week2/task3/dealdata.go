package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var uv_00 map[uint32]int8

var pv_00 map[string]int

var isp_00 map[string]int

var region_00 map[string]int

var visitStatus_00 map[string]int

type logInfoStru struct {
	Ip       uint32
	Hour     string
	PageName string
	Status   string
	ipInf    *ipInfo
}

type ipDataStore struct {
	dataStart uint32
	indexNums uint32
	ipStore   []uint32
	store     []byte
}

var vipDataStore ipDataStore

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Contains(f.Name(), ".log") {
			fmt.Println(f.Name())
			analysis(f.Name())
		}
		return nil

	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func analysis(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	r := bufio.NewReaderSize(f, 1024*1024*10)

	hour := "00"

	for {

		strbyte, _, _ := r.ReadLine()

		if strbyte == nil {
			AddData(filename, hour)
			fmt.Println("The End")
			fmt.Println(time.Now().String())
			break
		}

		str := string(strbyte)
		vlogInfoStru, ok := splitLog(str)
		if ok {
			if vlogInfoStru.Hour == hour {
				bufData(vlogInfoStru, uv_00, pv_00, isp_00, region_00, visitStatus_00)
			} else {
				AddData(filename, hour)

				hour = vlogInfoStru.Hour

				bufData(vlogInfoStru, uv_00, pv_00, isp_00, region_00, visitStatus_00)

			}
		}

	}

}

func bufData(vlogInfoStru *logInfoStru, uv map[uint32]int8, pv map[string]int, isp map[string]int, region map[string]int, visitStatus map[string]int) {

	//UV
	_, ok := uv[vlogInfoStru.Ip]
	if !ok {
		uv[vlogInfoStru.Ip] = 1
	}

	//PV
	_, okpv := pv["pv"]
	if okpv {
		pv["pv"] += 1
	} else {
		pv["pv"] = 1
	}

	//isp
	if vlogInfoStru.ipInf.country == "中国" {
		_, okisp := isp[vlogInfoStru.ipInf.isp]
		if okisp {
			isp[vlogInfoStru.ipInf.isp] += 1
		} else {
			isp[vlogInfoStru.ipInf.isp] = 1
		}
	} else {
		_, okisp := isp["other"]
		if okisp {
			isp["other"] += 1
		} else {
			isp["other"] = 1
		}
	}

	//region
	if vlogInfoStru.ipInf.country == "中国" {
		_, okregion := region[vlogInfoStru.ipInf.region]
		if okregion {
			region[vlogInfoStru.ipInf.region] += 1
		} else {
			region[vlogInfoStru.ipInf.region] = 1
		}
	}

	//visitStatus
	_, okvisitStatus := visitStatus[vlogInfoStru.Status]
	if okvisitStatus {
		visitStatus[vlogInfoStru.Status] += 1
	} else {
		visitStatus[vlogInfoStru.Status] = 1
	}
}

func splitLog(logContent string) (vlogInfoStru *logInfoStru, isValid bool) {

	isValid = true

	ipStrArr := strings.Split(logContent, " ")
	//fmt.Println(ipStrArr[0])

	timeStrArr := strings.Split(logContent, ":")
	//fmt.Println(timeStrArr[1])

	pageNameArr := strings.Split(logContent, "/")
	pageNameArr2 := strings.Split(pageNameArr[5], "\"")
	//fmt.Println(pageNameArr2[0])
	for _, v := range NotPV {
		if strings.Contains(pageNameArr2[0], v) {
			isValid = false
		}
	}

	ipUint32 := Ip2Uint32(ipStrArr[0])
	ipInf := ipQuery(ipUint32)

	if ipInf.isValid == 0 {
		isValid = false
	}

	//fmt.Println(ipStrArr[7])
	vlogInfoStru = &logInfoStru{ipUint32, timeStrArr[1], pageNameArr2[0], ipStrArr[7], ipInf}
	//vlogInfoStru = &vlogInfoStruA

	return
}

func makeMap() {
	uv_00 = make(map[uint32]int8)

	pv_00 = make(map[string]int)

	isp_00 = make(map[string]int)

	region_00 = make(map[string]int)

	visitStatus_00 = make(map[string]int)
}

func init() {

	vipDataStore = prepareIpData()
	makeMap()

}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Begin" + time.Now().String())
	getFilelist("./")
	//analysis("www.asta.com.log")
}
