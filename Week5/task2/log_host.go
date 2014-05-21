package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
)

func server() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer[:n]), addr)
	}
}

func client() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, _ := net.DialUDP("udp", nil, addr)

	defer conn.Close()
	for {
		var str string
		fmt.Scanf("%s", &str)
		io.WriteString(conn, str)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go server()
	go client()
	select {}
}
