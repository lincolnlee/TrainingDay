package main

import (
	"fmt"
	"html/template"
	//"log"
	"net/http"
	//"strings"
)

func setting(w http.ResponseWriter, r *http.Request) {

	var RssSourceLs *[]RssSource
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "POST" {
		fmt.Println(r.FormValue("rssName"), r.FormValue("rssUrl"))
		if r.FormValue("rssName") != "" && r.FormValue("rssUrl") != "" {
			AddSettingData(r.FormValue("rssName"), r.FormValue("rssUrl"))
		} else if r.FormValue("crawl") == "crawl" {
			//fmt.Println("crawl ok")
			go dealCrawl(RssSourceLs)
		}
	} else {
		if r.FormValue("action") == "del" {
			id := 0
			fmt.Sscanf(r.FormValue("id"), "%d", &id)
			DelSettingData(id)
		}
	}

	RssSourceLs = GetAllSettingData()
	//fmt.Println(RssSourceLs)
	t, err := template.ParseFiles("resource/setting.gtpl")
	checkError(err)
	t.Execute(w, RssSourceLs)
}

func dealCrawl(RssSourceLs *[]RssSource) {
	RssSourceLs = GetAllSettingData()
	ch := make(chan int, 2)
	for i, v := range *RssSourceLs {
		ch <- i
		go crawlRss(v.RssUrl, int32(v.Id), ch)
		fmt.Println(len(ch))
	}

	/*for {
		select {
		case c <- x:

		case <-quit:
			fmt.Println("quit")
			return
		}
	}*/
}
