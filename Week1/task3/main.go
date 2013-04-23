/**
 * Created with IntelliJ IDEA.
 * User: leelin
 * Date: 13-3-31
 * Time: 下午10:20
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Begin")

	//var myRegexp = regexp.MustCompile(`([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}`)

	f, _ := os.Open("test.log")

	defer f.Close()
	domainMap := make(map[string]*os.File)
	fmt.Println(time.Now().String())

	r := bufio.NewReaderSize(f, 1024*1024*10)

	for {

		strbyte, isPrefix, _ := r.ReadLine()
		if isPrefix {
			fmt.Println("deal 50MB time:" + time.Now().String())
			break
		}

		if strbyte == nil {
			fmt.Println("The End" + time.Now().String())
			return
		}

		str := string(strbyte)

		charIdx := 0
		startIdx := 0
		endIdx := 0
		for i, v := range str {
			if v == '/' {
				if startIdx > 0 {
					endIdx = i
					break
				} else if charIdx == i-1 {
					startIdx = i + 1
				} else {
					charIdx = i
				}

			}
		}
		rest := str[startIdx:endIdx]
		//rest := myRegexp.FindString(str)
		_, ok := domainMap[rest]
		if ok {

			domainMap[rest].WriteString(str)

		} else {
			fret, _ := os.Create(rest + ".log")
			domainMap[rest] = fret
			defer domainMap[rest].Close()
			domainMap[rest].WriteString(str)
		}
	}
}
