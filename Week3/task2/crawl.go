package main

import (
	"encoding/xml"
	"fmt"
	//"io"
	iconv "github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	//"strings"
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
	Item item `xml:"item"`
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

func main() {
	//con := rss{}
	ret, _ := http.Get("http://data.earthquake.cn/datashare/globeEarthquake_csn.html")
	defer ret.Body.Close()
	b, _ := ioutil.ReadAll(ret.Body)

	//fmt.Println(string(b))

	/*str := string(b)
	strArr := strings.Split(str, "?>")
	strArr2 := strings.Split(strArr[0], "'")
	fmt.Println(strArr2[3])

	if strings.ToUpper(strArr2[3]) != "UTF-8" {*/
	out := make([]byte, len(b))
	out = out[:]
	iconv.Convert(b, out, "GB2312", "UTF-8")
	b = out
	//}

	fmt.Println(string(b))

	//err := xml.Unmarshal(b, &con)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}

	//fmt.Println(con)
}
