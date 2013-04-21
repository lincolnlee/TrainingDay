package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"strings"
)

func defaultP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法

	pageIdx := 0
	fmt.Sscanf(r.FormValue("page"), "%d", &pageIdx)
	startIdx := 0
	if pageIdx > 1 {
		startIdx = (pageIdx - 1) * 20
	} else if pageIdx == 0 {
		pageIdx = 1
	}

	sumItemAmount := GetAllRssDataAmount()

	pageAmount := int(sumItemAmount) / 20
	if sumItemAmount%20 > 0 {
		pageAmount++
	}
	fmt.Println(pageAmount)
	fmt.Println(pageIdx)
	var page []PageInfo
	if pageAmount > 5 {
		page = make([]PageInfo, 5)

		if pageIdx < 4 {
			for i := 0; i < 5; i++ {
				page[i] = PageInfo{i+1 == pageIdx, i + 1}
			}
		} else if pageIdx <= pageAmount-2 {
			for i := 0; i < 5; i++ {
				page[i] = PageInfo{pageIdx+i-2 == pageIdx, pageIdx + i - 2}
			}
		} else {
			for i := 0; i < 5; i++ {
				page[i] = PageInfo{pageAmount+i-4 == pageIdx, pageAmount + i - 4}
			}
		}

	} else {
		page = make([]PageInfo, pageAmount)

		for i := 0; i < pageAmount; i++ {
			page[i] = PageInfo{i+1 == pageIdx, i + 1}
		}
	}

	RssContentLs := GetAllRssData(int32(startIdx), 20)

	RssContentPageD := RssContentPage{pageIdx, &page, RssContentLs}
	//fmt.Println(RssContentPageD.RssContentLs)
	t, err := template.ParseFiles("resource/default.gtpl")
	checkError(err)
	t.Execute(w, RssContentPageD)

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("resource/index.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", login)           //设置访问的路由
	http.HandleFunc("/index", login)      //设置访问的路由
	http.HandleFunc("/default", defaultP) //设置访问的路由
	http.HandleFunc("/setting", setting)  //设置访问的路由
	fileServer := http.FileServer(http.Dir("resource/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe("127.0.0.1:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
