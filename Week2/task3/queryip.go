package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type ipInfo struct {
	isValid int
	country string
	region  string
	city    string
	isp     string
}

func prepareIpData() ipDataStore {

	f, _ := os.Open("iploc.dat")
	defer f.Close()

	r := bufio.NewReaderSize(f, 1024*1024*4)

	cb, _ := r.Peek(12)

	indexStart := binary.LittleEndian.Uint32(cb[:4])
	dataStart := binary.LittleEndian.Uint32(cb[4:8])
	indexNums := binary.LittleEndian.Uint32(cb[8:12])

	store, _ := r.Peek(r.Buffered())

	indexStore := store[indexStart:]

	ipStore := make([]uint32, indexNums)

	for i := uint32(0); i < indexNums; i++ {
		ipStore[i] = binary.LittleEndian.Uint32(indexStore[i*4 : (i+1)*4])
	}

	return ipDataStore{dataStart, indexNums, ipStore, store}
}

func ipQuery(ipUint32 uint32) *ipInfo {

	low := uint32(0)
	high := vipDataStore.indexNums - 2
	middle := (low + high) / 2

	dataIndex := uint32(0)

	//ipUint32 := Ip2Uint32(ipStr)

	for low <= high {
		middle = (low + high) / 2

		if ipUint32 >= vipDataStore.ipStore[middle] && ipUint32 < vipDataStore.ipStore[middle+1] {
			dataIndex = middle*uint32(21) + vipDataStore.dataStart
			break
		} else if ipUint32 < vipDataStore.ipStore[middle] {
			high = middle - 1

		} else {
			low = middle + 1

		}
	}

	var ipInf ipInfo

	flag := vipDataStore.store[dataIndex : dataIndex+1]
	country := ""
	var region uint16
	var city uint32
	var isp uint16
	//var timezone uint16
	//var longitude, latitude float32

	if int(flag[0]) == 2 {
		if vipDataStore.store[dataIndex+1 : dataIndex+2][0] != byte(0x00) && vipDataStore.store[dataIndex+2 : dataIndex+3][0] != byte(0x00) {

			country = string(vipDataStore.store[dataIndex+1 : dataIndex+3])
			region = binary.LittleEndian.Uint16(vipDataStore.store[dataIndex+3 : dataIndex+5])
			city = binary.LittleEndian.Uint32(vipDataStore.store[dataIndex+5 : dataIndex+9])
			isp = binary.LittleEndian.Uint16(vipDataStore.store[dataIndex+9 : dataIndex+11])
			//timezone = binary.LittleEndian.Uint16(store[dataIndex+11 : dataIndex+13])
			//longitude = math.Float32frombits(binary.LittleEndian.Uint32(store[dataIndex+13 : dataIndex+17]))
			//latitude = math.Float32frombits(binary.LittleEndian.Uint32(store[dataIndex+17 : dataIndex+21]))

			var strregion []byte
			var strcity []byte
			var strisp []byte

			//fmt.Println("国家", COUNTRIES_ZH[country])

			if i := uint16(bytes.IndexByte(vipDataStore.store[region:], byte(0x00))); i >= 0 {
				strregion = vipDataStore.store[region : region+i+1]
				//fmt.Println("区域", string(strregion))
			}

			if i := uint32(bytes.IndexByte(vipDataStore.store[city:], byte(0x00))); i >= 0 {
				strcity = vipDataStore.store[city : city+i+1]
				//fmt.Println("城市", string(strcity))
			}

			if i := uint16(bytes.IndexByte(vipDataStore.store[isp:], byte(0x00))); i >= 0 {
				strisp = vipDataStore.store[isp : isp+i+1]
				//fmt.Println("运营商", string(strisp))
			}

			//if i := uint16(bytes.IndexByte(store[timezone:], byte(0x00))); i >= 0 {
			//	strtimezone := store[timezone : timezone+i+1]
			//	fmt.Println("时区", string(strtimezone))
			//}

			//fmt.Println("经度", longitude)
			//fmt.Println("纬度", latitude)

			ipInf = ipInfo{1, COUNTRIES_ZH[country], string(strregion), string(strcity), string(strisp)}
		} else {
			//fmt.Println("No")
			ipInf = ipInfo{0, "", "", "", ""}
		}
	} else if int(flag[0]) == 1 {
		ipInf = ipInfo{0, "", "", "", ""}
		//fmt.Println("IANA保留地址")
	} else {
		ipInf = ipInfo{0, "", "", "", ""}
		//fmt.Println("IANA未分配地址")
	}

	return &ipInf

}

func Ip2Uint32(str string) uint32 {
	var a, b, c, d byte
	n, err := fmt.Sscanf(str, "%d.%d.%d.%d", &a, &b, &c, &d)
	if err != nil || n != 4 {
		return 0
	}
	ip := uint32(a) << 24
	ip |= uint32(b) << 16
	ip |= uint32(c) << 8
	ip |= uint32(d)
	return ip
}
