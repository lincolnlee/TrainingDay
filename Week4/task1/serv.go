package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	service := ":8001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go Calculate(conn)
	}
}

func Calculate(rwc net.Conn) {
	headData := make([]byte, 12)
	var length int64 = 2060

	c, err := rwc.Read(headData)
	if err == nil {
		length, _ = strconv.ParseInt(string(headData[0:c]), 10, 64)
	} else {
		rwc.Write([]byte(err.Error()))
		rwc.Close()
		return
	}
	data := make([]byte, length)
	c, err = rwc.Read(data)
	if err == nil {
		answer := dealCalculate(string(data[0:c]))
		rwc.Write([]byte(strconv.FormatFloat(answer, 'f', -1, 64)))
	} else {
		rwc.Write([]byte(err.Error()))
		rwc.Close()
		return
	}

	rwc.Close()
}

func dealCalculate(exper string) (ret float64) {
	mExprSlice := strings.Split(exper, "|")
	OperandSlice := make([]float64, 100)
	i := -1

	for _, v := range mExprSlice {
		if IsNum(v) {
			i++
			OperandSlice[i], _ = strconv.ParseFloat(v, 64)
		} else {
			switch v {
			case "+":
				OperandSlice[i-1] = OperandSlice[i] + OperandSlice[i-1]
				i--
			case "-":
				OperandSlice[i-1] = OperandSlice[i] - OperandSlice[i-1]
				i--
			case "*":
				OperandSlice[i-1] = OperandSlice[i] * OperandSlice[i-1]
				i--
			case "/":
				OperandSlice[i-1] = OperandSlice[i] / OperandSlice[i-1]
				i--
			}
		}
	}
	ret = OperandSlice[0]
	return
}

func IsNum(s string) (ret bool) {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		ret = false
	} else {
		ret = true
	}
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
