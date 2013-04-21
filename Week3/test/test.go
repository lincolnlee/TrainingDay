package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"rss"`
	//Version string   `xml:"itunes,attr"`
	//b       string   `xml:"dc,attr"`
	//c       string   `xml:"taxo,attr"`
	//d       string   `xml:"rdf,attr"`
	Svs []server `xml:"channel"`
	//Description string `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"channel"`
	ServerName string   `xml:"title"`
	ServerIP   string   `xml:"link"`
	ServerIP2  string   `xml:"description"`
	//description   string   `xml:"description"`
	Language2     string `xml:"language"`
	PubDate2      string `xml:"pubDate"`
	LastBuildDate string `xml:"lastBuildDate"`
	Ttl           string `xml:"ttl"`
}

func main() {
	file, err := os.Open("servers.xml") // For read access.     
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
