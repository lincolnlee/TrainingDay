package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//机器类，看了《I o P》，对这个名字有莫名其妙的好感哈。
//主要的收发命令的业务逻辑都在这个类中实现
type Machine struct {
	AppConfig *AppConfig
}

func (this *Machine) StartHTTPService() {
	http.HandleFunc("/", this.DealReceiveCmd)                                        //设置访问的路由
	go http.ListenAndServe(this.AppConfig.Me.Ip+":"+this.AppConfig.Me.HttpPort, nil) //设置监听的端口
}

func (this *Machine) StartTCPService() {
	service := this.AppConfig.Me.Ip + ":" + this.AppConfig.Me.TCPPort
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	utility.checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	utility.checkError(err)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go this.SendFile(conn)
		}
	}()
}

func (this *Machine) SendFile(rwc net.Conn) {
	defer rwc.Close()

	f, _ := os.Open("iploc.dat")
	defer f.Close()

	fi, _ := f.Stat()
	fileSize := fi.Size()

	_, err := rwc.Write([]byte(utility.formatLength(fileSize)))
	utility.checkError(err)

	r := bufio.NewReaderSize(f, 1024*1024*1)

	p := make([]byte, 1024*1024*1)

	for {
		byteAmount, err := r.Read(p)

		if err == io.EOF {
			break
		} else if err != nil {
			utility.checkError(err)
		}

		fmt.Println(byteAmount)
		fmt.Println(p[(byteAmount - 64):byteAmount])

		_, err = rwc.Write(p[0:byteAmount])
		utility.checkError(err)
	}
}

func (this *Machine) DealSendCmd() {
	strCmd := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strCmd = scanner.Text()
		if strCmd == ":q" {
			os.Exit(1)
		} else {

			command := Command{Source: this.AppConfig.Me.Ip + ":" + this.AppConfig.Me.TCPPort, Action: strCmd}
			strJson, _ := json.Marshal(command)
			fmt.Println(string(strJson))
			for _, v := range this.AppConfig.Hosts {

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

func (this *Machine) DealReceiveCmd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	//fmt.Println(r)
	if r.Method == "GET" {
		strCmd, _ := url.QueryUnescape(r.FormValue("cmd"))
		fmt.Println("cmd:", strCmd)

		var cmd Command
		json.Unmarshal([]byte(strCmd), &cmd)

		if !this.CustomCmd(cmd) {
			cmdObj := exec.Command(cmd.Action)

			out, err := cmdObj.Output()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(out))
		}
	}
}

func (this *Machine) CustomCmd(cmd Command) bool {
	if strings.Contains(cmd.Action, ":") {
		if cmd.Action == ":getfile" {
			go func() {
				service := cmd.Source

				tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
				utility.checkError(err)

				conn, err := net.DialTCP("tcp4", nil, tcpAddr)
				utility.checkError(err)

				headData := make([]byte, 20)
				var length int64 = 2060

				c, err := conn.Read(headData)
				utility.checkError(err)

				length, err = strconv.ParseInt(string(headData[0:c]), 10, 64)
				utility.checkError(err)
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
						utility.checkError(err)
					}

					fmt.Println("read size:", c)

					receiveLength += int64(c)

					fmt.Println("sum size:", receiveLength)
					fmt.Println("end :", data[c-64:c])

					fret.Write(data[0:c])
				}

				conn.Close()
			}()
		}
		return true
	} else {
		return false
	}
}
