package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	//"net"
	"net/http"
	"os"
	//"time"
	"bufio"
	//"bytes"
	"io"
	"net"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

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

type Command struct {
	Source string
	Action string
	From   string
	To     string
}

func main() {
	fmt.Println("Start...")
	appConfig := GetCfg()
	StartHTTPService(appConfig)
	StartTCPService(appConfig)
	DealSendCmd(appConfig)
}

func GetCfg() AppConfig {
	appConfig := AppConfig{}

	file, err := os.Open("app.xml")
	checkError(err)

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	checkError(err)

	xml.Unmarshal(data, &appConfig)
	fmt.Println(appConfig)

	return appConfig
}

func StartHTTPService(appConfig AppConfig) {
	http.HandleFunc("/", DealReceiveCmd)                                   //设置访问的路由
	go http.ListenAndServe(appConfig.Me.Ip+":"+appConfig.Me.HttpPort, nil) //设置监听的端口
}

func StartTCPService(appConfig AppConfig) {
	service := appConfig.Me.Ip + ":" + appConfig.Me.TCPPort
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go SendFile(conn)
		}
	}()
}

func SendFile(rwc net.Conn) {
	defer rwc.Close()

	f, _ := os.Open("iploc.dat")
	defer f.Close()

	fi, _ := f.Stat()
	fileSize := fi.Size()

	_, err := rwc.Write([]byte(formatLength(fileSize)))
	checkError(err)

	r := bufio.NewReaderSize(f, 1024*1024*1)

	p := make([]byte, 1024*1024*1)

	for {
		byteAmount, err := r.Read(p)

		if err == io.EOF {
			break
		} else if err != nil {
			checkError(err)
		}

		fmt.Println(byteAmount)
		fmt.Println(p[(byteAmount-64):byteAmount])

		_, err = rwc.Write(p[0:byteAmount])
		checkError(err)
	}

	/*var sendLength int64 = 0
	for {
		p, err := r.Peek(1024 * 1024 * 1)
		checkError(err)

		sendLength += int64(len(p))
		fmt.Println(len(p))

		_, err = rwc.Write(p)
		checkError(err)

		if sendLength >= fileSize {
			break
		}
	}*/
}

func DealSendCmd(appConfig AppConfig) {
	strCmd := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strCmd = scanner.Text()
		if strCmd == ":q" {
			os.Exit(1)
		} else {

			command := Command{Source: appConfig.Me.Ip + ":" + appConfig.Me.TCPPort, Action: strCmd}
			strJson, _ := json.Marshal(command)
			fmt.Println(string(strJson))
			for _, v := range appConfig.Hosts {

				response, err := http.Get("http://" + v.Ip + ":" + v.HttpPort + "/?cmd=" + url.QueryEscape(string(strJson)))
				if err != nil {
					fmt.Println("err:" + err.Error())
				} else {
					fmt.Println(response.Body)
				}

			}
			fmt.Println("Input cmd.")
		}
	}

	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "reading standard input:", err)

	}
}

func DealReceiveCmd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	//fmt.Println(r)
	if r.Method == "GET" {
		strCmd, _ := url.QueryUnescape(r.FormValue("cmd"))
		fmt.Println("cmd:", strCmd)

		var cmd Command
		json.Unmarshal([]byte(strCmd), &cmd)

		if !CustomCmd(cmd) {
			cmdObj := exec.Command(cmd.Action)

			out, err := cmdObj.Output()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(out))
		}
	}
}

func CustomCmd(cmd Command) bool {
	if strings.Contains(cmd.Action, ":") {
		if cmd.Action == ":getfile" {
			go func() {
				service := cmd.Source

				tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
				checkError(err)

				conn, err := net.DialTCP("tcp4", nil, tcpAddr)
				checkError(err)

				headData := make([]byte, 20)
				var length int64 = 2060

				c, err := conn.Read(headData)
				checkError(err)

				length, err = strconv.ParseInt(string(headData[0:c]), 10, 64)
				checkError(err)
				fmt.Println("file size:", length)

				fret, _ := os.Create("iploc.dat")
				defer fret.Close()

				data := make([]byte, length)

				var receiveLength int64 = 0
				for {
					c, err = conn.Read(data)
					if err == io.EOF {
						break
					} else if err != nil {
						checkError(err)
					}

					fmt.Println("read size:", c)

					receiveLength += int64(c)

					fmt.Println("sum size:", receiveLength)
					fmt.Println("end :", data[c-64:c])

					fret.Write(data[0:c])
				}

				//ret := make([]byte, 1024*1024*8)
				//conn.Read(ret)
				//ret, err := ioutil.ReadAll(conn)
				//checkError(err)

				//fmt.Println(ret)

				conn.Close()
			}()
		}
		return true
	} else {
		return false
	}
}
