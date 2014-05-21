package main

import (
	"bufio"
	//"bytes"
	//"encoding/binary"
	"fmt"
	//"math"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	NotPV     []string = []string{".css", ".js", ".class", ".gif", ".jpg", ".jpeg", ".png", ".bmp", ".ico", "rss", "xml", "swf"}
	rest      string   = "test"
	domainMap          = make(map[string]*os.File)

	daAsMap = make(map[string]*asStru)
	hoAsMap = make(map[string]*asStru)

	dirName string = ""
)

type asStru struct {
	icou   int
	scou   int
	allcou int
}

func getFilelist(path string) {
	_, ok := domainMap[rest]

	if !ok {
		fo, openFileErr := os.OpenFile(rest+".log", os.O_RDWR|os.O_APPEND, 0660)
		if openFileErr != nil && os.IsNotExist(openFileErr) {
			fret, createErr := os.Create(rest + ".log")
			if createErr == nil {
				defer fret.Close()
				fo, openFileErr := os.OpenFile(rest+".log", os.O_RDWR|os.O_APPEND, 0660)
				if openFileErr == nil {
					defer fo.Close()
					domainMap[rest] = fo
				}
			}
		} else if openFileErr == nil {
			defer fo.Close()
			domainMap[rest] = fo
		}

	}

	var wg sync.WaitGroup

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			fmt.Println(f.Name())
			dirName = f.Name()
			return nil
		}
		if strings.Contains(f.Name(), ".log") && f.Name() != (rest+".log") {
			fmt.Println(f.Name())
			daylogname := dirName + f.Name()[0:10]
			hourelogname := dirName + f.Name()
			if daAsMap[daylogname] == nil {
				daAsMap[daylogname] = &asStru{0, 0, 0}
			}
			if hoAsMap[hourelogname] == nil {
				hoAsMap[hourelogname] = &asStru{0, 0, 0}
			}
			wg.Add(1)
			go analysis(path, hourelogname, &wg)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	wg.Wait()

	for k, v := range hoAsMap {
		if daAsMap[k[0:13]] == nil {
			daAsMap[k[0:13]] = &asStru{0, 0, 0}
		}

		daAsMap[k[0:13]].icou = daAsMap[k[0:13]].icou + v.icou
		daAsMap[k[0:13]].scou = daAsMap[k[0:13]].scou + v.scou
		daAsMap[k[0:13]].allcou = daAsMap[k[0:13]].allcou + v.allcou
	}

	for k, v := range daAsMap {
		fmt.Println(k)
		fmt.Println("timeout:", v.icou)
		fmt.Println("badresponse:", v.scou)
		fmt.Println("allresponse:", v.allcou)
		_, ok := domainMap[rest]

		if ok {
			domainMap[rest].WriteString(k + "\n")
			domainMap[rest].WriteString("timeout:" + strconv.Itoa(v.icou) + "\n")
			domainMap[rest].WriteString("badresponse:" + strconv.Itoa(v.scou) + "\n")
			domainMap[rest].WriteString("allresponse:" + strconv.Itoa(v.allcou) + "\n")
		}
	}
}

func analysis(filename string, hourelogname string, wg *sync.WaitGroup) {
	f, _ := os.Open(filename)

	defer f.Close()

	r := bufio.NewReaderSize(f, 1024*1024*10)

	for {

		strbyte, _, _ := r.ReadLine()

		if strbyte == nil {
			//fmt.Println("The End" + time.Now().String())
			break
		}

		str := string(strbyte)
		if string(str[0]) != "#" {
			splitLog(str, hourelogname)
		}

	}
	wg.Done()
	//runtime.GC()

}

func splitLog(logContent string, hourelogname string) {

	ipStrArr := strings.Split(logContent, " ")
	//fmt.Println(ipStrArr[(len(ipStrArr) - 2)])
	time_taken, _ := strconv.Atoi(ipStrArr[len(ipStrArr)-1])
	sc_status, _ := strconv.Atoi(ipStrArr[len(ipStrArr)-4])

	hoAsMap[hourelogname].allcou = hoAsMap[hourelogname].allcou + 1

	if time_taken > 5000 {
		//fmt.PrintLn(logContent)
		hoAsMap[hourelogname].icou = hoAsMap[hourelogname].icou + 1

		//_, ok := domainMap[rest]

		//if ok {
		//	domainMap[rest].WriteString(logContent + "\n")
		//}
	}

	if sc_status > 399 {
		//fmt.PrintLn(logContent)
		hoAsMap[hourelogname].scou = hoAsMap[hourelogname].scou + 1

		//_, ok := domainMap[rest]

		//if ok {
		//	domainMap[rest].WriteString(logContent + "\n")
		//}
	}

	//return
}

func init() {

}

func main() {
	runtime.GOMAXPROCS(16)
	fmt.Println("Begin" + time.Now().String())
	getFilelist("./")
	//analysis("u_ex14021620.log")
	fmt.Println("The End" + time.Now().String())
}
