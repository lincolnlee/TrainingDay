package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	//"math"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var COUNTRIES_ZH = map[string]string{
	"AP": "亚洲/太平洋地区",
	"EU": "欧洲",
	"AD": "安道​​尔",
	"AE": "阿拉伯联合大公国",
	"AF": "阿富汗",
	"AG": "安提瓜和巴布达",
	"AI": "安圭拉",
	"AL": "阿尔巴尼亚",
	"AM": "亚美尼亚",
	"AN": "荷属安的列斯群岛",
	"AO": "安哥拉",
	"AQ": "南极洲",
	"AR": "阿根廷",
	"AS": "美属萨摩亚",
	"AT": "奥地利",
	"AU": "澳大利亚",
	"AW": "阿鲁巴",
	"AZ": "阿塞拜疆",
	"BA": "波斯尼亚和黑塞哥维那",
	"BB": "巴巴多斯",
	"BD": "孟加拉国",
	"BE": "比利时",
	"BF": "布基纳法索",
	"BG": "保加利亚",
	"BH": "巴林",
	"BI": "布隆迪",
	"BJ": "贝宁",
	"BM": "百慕大",
	"BN": "文莱",
	"BO": "玻利维亚",
	"BR": "巴西",
	"BS": "巴哈马",
	"BT": "不丹",
	"BV": "布维岛",
	"BW": "博茨瓦纳",
	"BY": "白俄罗斯",
	"BZ": "伯利兹",
	"CA": "加拿大",
	"CC": "科科斯群岛",
	"CD": "刚果民主共和国",
	"CF": "中非共和国",
	"CG": "刚果",
	"CH": "瑞士",
	"CI": "科特迪瓦",
	"CK": "库克群岛",
	"CL": "智利",
	"CM": "喀麦隆",
	"CN": "中国",
	"CO": "哥伦比亚",
	"CR": "哥斯达黎加",
	"CU": "古巴",
	"CV": "佛得角",
	"CX": "圣诞岛",
	"CY": "塞浦路斯",
	"CZ": "捷克共和国",
	"DE": "德国",
	"DJ": "吉布提",
	"DK": "丹麦",
	"DM": "多米尼加",
	"DO": "多明尼加共和国",
	"DZ": "阿尔及利亚",
	"EC": "厄瓜多尔",
	"EE": "爱沙尼亚",
	"EG": "埃及",
	"EH": "西撒哈拉",
	"ER": "厄立特里亚",
	"ES": "西班牙",
	"ET": "埃塞俄比亚",
	"FI": "芬兰",
	"FJ": "斐",
	"FK": "福克兰群岛（马尔维纳斯群岛）",
	"FM": "密克罗尼西亚联邦",
	"FO": "法罗群岛",
	"FR": "法国",
	"FX": "法国，大都会",
	"GA": "加蓬",
	"GB": "联合王国",
	"GD": "格林纳达",
	"GE": "格鲁吉亚",
	"GF": "法属圭亚那",
	"GH": "加纳",
	"GI": "直布罗陀",
	"GL": "格陵兰",
	"GM": "冈比亚",
	"GN": "几内亚",
	"GP": "瓜德罗普岛",
	"GQ": "赤道几内亚",
	"GR": "希腊",
	"GS": "南乔治亚岛和南桑威奇群岛",
	"GT": "危地马拉",
	"GU": "关岛",
	"GW": "几内亚比绍",
	"GY": "圭亚那",
	"HK": "香港",
	"HM": "赫德岛和麦克唐纳群岛",
	"HN": "洪都拉斯",
	"HR": "克罗地亚",
	"HT": "海地",
	"HU": "匈牙利",
	"ID": "印尼",
	"IE": "爱尔兰",
	"IL": "以色列",
	"IN": "印度",
	"IO": "英属印度洋领地",
	"IQ": "伊拉克",
	"IR": "伊朗伊斯兰共和国",
	"IS": "冰岛",
	"IT": "意大利",
	"JM": "牙买加",
	"JO": "约旦",
	"JP": "日本",
	"KE": "肯尼亚",
	"KG": "吉尔吉斯斯坦",
	"KH": "柬埔寨",
	"KI": "基里巴斯",
	"KM": "科摩罗",
	"KN": "圣基茨和尼维斯",
	"KP": "朝鲜人民民主共和国",
	"KR": "大韩民国",
	"KW": "科威特",
	"KY": "开曼群岛",
	"KZ": "哈萨克斯坦",
	"LA": "老挝人民民主共和国",
	"LB": "黎巴嫩",
	"LC": "圣卢西亚",
	"LI": "列支敦士登",
	"LK": "斯里兰卡",
	"LR": "利比里亚",
	"LS": "莱索托",
	"LT": "立陶宛",
	"LU": "卢森堡",
	"LV": "拉脱维亚",
	"LY": "利比亚",
	"MA": "摩洛哥",
	"MC": "摩纳哥",
	"MD": "摩尔多瓦共和国",
	"MG": "马达加斯加",
	"MH": "马绍尔群岛",
	"MK": "马其顿",
	"ML": "马里",
	"MM": "缅甸",
	"MN": "蒙古",
	"MO": "澳门",
	"MP": "北马里亚纳群岛",
	"MQ": "马提尼克",
	"MR": "毛里塔尼亚",
	"MS": "蒙特塞拉特",
	"MT": "马耳他",
	"MU": "毛里求斯",
	"MV": "马尔代夫",
	"MW": "马拉维",
	"MX": "墨西哥",
	"MY": "马来西亚",
	"MZ": "莫桑比克",
	"NA": "纳米比亚",
	"NC": "新喀里多尼亚",
	"NE": "尼日尔",
	"NF": "诺福克岛",
	"NG": "尼日利亚",
	"NI": "尼加拉瓜",
	"NL": "荷兰",
	"NO": "挪威",
	"NP": "尼泊尔",
	"NR": "瑙鲁",
	"NU": "纽埃",
	"NZ": "新西兰",
	"OM": "阿曼",
	"PA": "巴拿马",
	"PE": "秘鲁",
	"PF": "法属波利尼西亚",
	"PG": "巴布亚新几内亚",
	"PH": "菲律宾",
	"PK": "巴基斯坦",
	"PL": "波兰",
	"PM": "圣皮埃尔和密克隆",
	"PN": "皮特凯恩群岛",
	"PR": "波多黎各",
	"PS": "巴勒斯坦领土",
	"PT": "葡萄牙",
	"PW": "帕劳",
	"PY": "巴拉圭",
	"QA": "卡塔尔",
	"RE": "团圆",
	"RO": "罗马尼亚",
	"RU": "俄罗斯联邦",
	"RW": "卢旺达",
	"SA": "沙特阿拉伯",
	"SB": "所罗门群岛",
	"SC": "塞舌尔",
	"SD": "苏丹",
	"SE": "瑞典",
	"SG": "新加坡",
	"SH": "圣赫勒拿",
	"SI": "斯洛文尼亚",
	"SJ": "斯瓦尔巴群岛和扬马延岛",
	"SK": "斯洛伐克",
	"SL": "塞拉利昂",
	"SM": "圣马力诺",
	"SN": "塞内加尔",
	"SO": "索马里",
	"SR": "苏里南",
	"ST": "圣多美和普林西比",
	"SV": "萨尔瓦多",
	"SY": "阿拉伯叙利亚共和国",
	"SZ": "斯威士兰",
	"TC": "特克斯和凯科斯群岛",
	"TD": "乍得",
	"TF": "法国南部领土",
	"TG": "多哥",
	"TH": "泰国",
	"TJ": "塔吉克斯坦",
	"TK": "托克劳",
	"TM": "土库曼斯坦",
	"TN": "突尼斯",
	"TO": "汤加",
	"TL": "东帝汶",
	"TR": "土耳其",
	"TT": "特里尼达和多巴哥",
	"TV": "图瓦卢",
	"TW": "台湾",
	"TZ": "坦桑尼亚联合共和国",
	"UA": "乌克兰",
	"UG": "乌干达",
	"UM": "美国本土外小岛屿",
	"US": "美国",
	"UY": "乌拉圭",
	"UZ": "乌兹别克斯坦",
	"VA": "罗马教廷（梵蒂冈城国）",
	"VC": "圣文森特和格林纳丁斯",
	"VE": "委内瑞拉",
	"VG": "英属维尔京群岛",
	"VI": "美属维京群岛，",
	"VN": "越南",
	"VU": "瓦努阿图",
	"WF": "瓦利斯群岛和富图纳群岛",
	"WS": "萨摩亚",
	"YE": "也门",
	"YT": "马约特岛",
	"RS": "塞尔维亚",
	"ZA": "南非",
	"ZM": "赞比亚",
	"ME": "黑山",
	"ZW": "津巴布韦",
	"A1": "匿名代理",
	"A2": "卫星供应商",
	"O1": "其他",
	"AX": "奥兰群岛",
	"GG": "根西岛",
	"IM": "马恩岛",
	"JE": "新泽西州",
	"BL": "圣巴泰勒米",
	"MF": "圣马丁",
	"BQ": "博内尔，圣尤斯特歇斯和萨巴",
	"SS": "南苏丹",
}

var uv_00 map[uint32]int8

var pv_00 map[string]int

var isp_00 map[string]int

var region_00 map[string]int

var visitStatus_00 map[string]int

type ipInfo struct {
	isValid int
	country string
	region  string
	city    string
	isp     string
}

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

var (
	NotPV []string = []string{".css", ".js", ".class", ".gif", ".jpg", ".jpeg", ".png", ".bmp", ".ico", "rss", "xml", "swf"}
)

var vipDataStore ipDataStore

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
			fmt.Println("The End" + time.Now().String())
			break
		}

		str := string(strbyte)
		vlogInfoStru, ok := splitLog(str)
		if ok {
			if vlogInfoStru.Hour == hour {
				bufData(vlogInfoStru, uv_00, pv_00, isp_00, region_00, visitStatus_00)
			} else {
				fmt.Println("uv:", len(uv_00), "pv:", pv_00["pv"])
				fmt.Println("isp:", isp_00)
				fmt.Println("region:", region_00)
				fmt.Println("visitStatus:", visitStatus_00)

				db, _ := sql.Open("mysql", "root:ubuntu@tcp(127.0.0.1:3306)/deallog?charset=utf8")
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
				res, _ := stmt.Exec(strings.TrimRight(filename, ".log"), "2013-3-13", "00", len(uv_00), pv_00["pv"], ispS, regionS, visitStatusS)
				//checkErr(err)
				id, _ := res.LastInsertId()
				//checkErr(err)
				fmt.Println(id)

				db.Close()

				makeMap()

				runtime.GC()

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
