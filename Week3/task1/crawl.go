package main

import (
	"encoding/xml"
	"fmt"
	//"io"
	iconv "github.com/djimenez/iconv-go"
	"hash/adler32"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type image struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Link  string `xml:"link"`
}

type item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Caregory    string `xml:"caregory"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
}

type channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Language      string   `xml:"language"`
	PubDate       string   `xml:"pubDate"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Ttl           string   `xml:"ttl"`
	//Image         []image  `xml:"image"`
	Item []item `xml:"item"`
}

type rss struct {
	XMLName xml.Name `xml:"rss"`
	//A       string    `xml:"itunes,attr"`
	//B       string    `xml:"dc,attr"`
	//C       string    `xml:"taxo,attr"`
	//D       string    `xml:"rdf,attr"`
	Channel []channel `xml:"channel"`
	//Description string    `xml:",innerxml"`
}

func crawlRss(rurl string, rssid int32, ch chan int) {

	defer func() {
		if x := recover(); x != nil {
			<-ch
			fmt.Println(rssid, "panic")
			return
		} else {
			<-ch
			fmt.Println(rssid, "ok")
		}
	}()
	con := rss{}
	var ret *http.Response
	var err1 error

	for i := 0; i < 6; i++ {
		if i == 5 {
			<-ch
			return
		}

		//retry 5 times ,if fail more then 5 times return
		ret, err1 = http.Get(rurl)
		if err1 != nil {
			time.Sleep(5000 * time.Millisecond)
		} else {
			break
		}
	}

	defer ret.Body.Close()
	b, _ := ioutil.ReadAll(ret.Body)

	str := string(b)
	strArr := strings.Split(str, "?>")
	strArr2 := strings.Split(strArr[0], "\"")
	if len(strArr2) < 4 {
		strArr2 = strings.Split(strArr[0], "'")
	}
	//fmt.Println(strArr2[3])

	if strings.ToUpper(strArr2[3]) != "UTF-8" {
		out := make([]byte, len(b))
		out = out[:]
		iconv.Convert(b, out, strArr2[3], "utf-8")
		b = out
	}

	err := xml.Unmarshal(b, &con)
	checkError(err)

	if len(con.Channel) > 0 {
		if len(con.Channel[0].Item) > 0 {
			for _, v := range con.Channel[0].Item {
				hash1 := adler32.Checksum([]byte(v.Title))
				hash2 := adler32.Checksum([]byte(v.Link))

				amount := GetRssDataAmount(hash1, hash2)

				if amount == 0 {

					AddRssData(&RssContent{Atitle: v.Title, Adesc: v.Description, Alink: v.Link, Atime: ParseTime(v.PubDate), Srcid: rssid, Ahash1: hash1, Ahash2: hash2})
				}
			}
		}
	}

	//fmt.Println(con)
}

func ParseTime(timestr string) time.Time {
	formats := []string{"02/Jan/2006:15:04:05", "2006-01-02 15:04:05", time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano, time.Kitchen, time.Stamp, time.StampMilli}
	var t time.Time

	for _, format := range formats {
		t, err := time.Parse(format, timestr)
		if err != nil {
			//fmt.Println(err)
			continue
		}
		return t
	}
	return t
}

var conversion = map[string]string{
	"B": "January", //月 英文 完整
	"b": "Jan",     //月 英文 缩写
	"m": "01",      //月 数字
	"A": "Monday",  //周 英文 完整
	"a": "Mon",     //周 缩写 完整
	"d": "02",      //日 数字
	"H": "15",      //时 24小时制 数字
	"I": "03",      //时 12小时制 数字
	"M": "04",      //分 数字
	"S": "05",      //秒 数字
	"Y": "2006",    //年 完整 数字
	"y": "06",      //年 缩写 数字
	"p": "PM",      //12小时制 上下午 AM PM
	"Z": "MST",     //时区
	"z": "-0700",   //时区 数字
}

// Go的时间格式好坑爹，还是按Python的来吧！！
func Format(format string, t time.Time) string {
	layout := layoutParser(format)
	return t.Format(layout)
}

// Go的时间格式好坑爹，还是按Python的来吧！！
func Parse(format, value string) (time.Time, error) {
	layout := layoutParser(format)
	return time.Parse(layout, value)
}

func layoutParser(format string) string {
	formatChunks := strings.Split(format, "%")
	var layout []string
	for _, chunk := range formatChunks {
		if len(chunk) == 0 {
			continue
		}
		if layoutCmd, ok := conversion[chunk[0:1]]; ok {
			layout = append(layout, layoutCmd)
			if len(chunk) > 1 {
				layout = append(layout, chunk[1:])
			}
		} else {
			layout = append(layout, "%", chunk)
		}
	}
	return strings.Join(layout, "")
}
