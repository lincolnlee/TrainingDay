package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//配置类，负责读取配置文件

type Me struct {
	XMLName  xml.Name `xml:"me"`
	Ip       string   `xml:"ip,attr"`
	HttpPort string   `xml:"HttpPort,attr"`
	TCPPort  string   `xml:"TCPPort,attr"`
}

type Host struct {
	XMLName  xml.Name `xml:"host"`
	Ip       string   `xml:"ip,attr"`
	HttpPort string   `xml:"HttpPort,attr"`
	TCPPort  string   `xml:"TCPPort,attr"`
}

type AppConfig struct {
	XMLName xml.Name `xml:"root"`
	Me      Me       `xml:"me"`
	Hosts   []Host   `xml:"host"`
}

func (this *AppConfig) GetCfg() {

	file, err := os.Open("app.xml")
	utility.checkError(err)

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	utility.checkError(err)

	xml.Unmarshal(data, this)
	fmt.Println(this)
}
