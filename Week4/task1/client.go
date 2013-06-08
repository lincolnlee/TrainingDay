package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	//"reflect"
	"strconv"
	"strings"
)

var strOperators1 string = "+-"
var strOperators2 string = "*/a"
var strOperators0 string = "()"
var strLeftParenthesis string = "("
var strRightParenthesis string = ")"

func main() {
	mExpr1 := "3 + 5 * 7 + 9 * 11 + 13"
	//mExpr2 := "3 + 5 * ( 7 + 9 ) * 11 + 13"
	//mExpr3 := "3 + 5 + 7 + 9 * 11 + 13"
	//mExpr4 := "3 + 5 * ( 7 + 9 * 11 ) + 13"
	l1 := Convert(mExpr1)
	l1 = strings.TrimRight(l1, "|")
	length := len(l1)
	strLen := formatLength(length)

	l1 = strLen + l1

	fmt.Println(strLen)
	//l2 := Convert(mExpr2)
	//l3 := Convert(mExpr3)
	//l4 := Convert(mExpr4)

	//fmt.Println(mExpr1)
	//fmt.Println(l1)

	//fmt.Println(mExpr2)
	//fmt.Println(l2)

	//fmt.Println(mExpr3)
	//fmt.Println(l3)

	//fmt.Println(mExpr4)
	//fmt.Println(l4)

	//if len(os.Args) != 2 {
	//	fmt.Println(os.Stderr, "Usage: %s host:port", os.Args[0])
	//	os.Exit(1)
	//}
	//fmt.Println(os.Args)
	service := ":8001" //os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte(l1))
	checkError(err)

	ret, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(ret))

	conn.Close()
	os.Exit(1)
}

func CompareOperators(ope1 string, ope2 string) (ret int) {
	opeLv1 := 1
	opeLv2 := 1

	if strings.Contains(strOperators1, ope1) {
		opeLv1 = 1
	}

	if strings.Contains(strOperators2, ope1) {
		opeLv1 = 2
	}

	if strings.Contains(strOperators0, ope1) {
		opeLv1 = 0
	}

	if strings.Contains(strOperators1, ope2) {
		opeLv2 = 1
	}

	if strings.Contains(strOperators2, ope2) {
		opeLv2 = 2
	}

	if strings.Contains(strOperators0, ope2) {
		opeLv2 = 0
	}

	if opeLv1 > opeLv2 {
		ret = 2
	} else if opeLv1 == opeLv2 {
		ret = 1
	} else {
		ret = 0
	}
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

func Convert(mExpr string) (ret string) {
	mExprSlice := strings.Split(mExpr, " ")

	//numSlice := make([]string, 10)
	OperatorsSlice := make([]string, 100)

	i := -1

	for _, v := range mExprSlice {
		//fmt.Println(v)
		if IsNum(v) {
			//fmt.Println("a")
			ret += v + "|"
		} else if strings.Contains(strLeftParenthesis, v) {
			//fmt.Println("b")
			i++
			OperatorsSlice[i] = v
		} else if strings.Contains(strRightParenthesis, v) {
			//fmt.Println("c")
			for j := i; j >= 0; j-- {
				if !strings.Contains(strLeftParenthesis, OperatorsSlice[j]) {
					//fmt.Println(OperatorsSlice[j])
					ret += OperatorsSlice[j] + "|"
					i--
				} else {
					i--
					j = 0
				}
			}
		} else if i >= 0 && CompareOperators(v, OperatorsSlice[i]) == 2 {
			//fmt.Println("d")
			i++
			OperatorsSlice[i] = v
		} else if i >= 0 && CompareOperators(v, OperatorsSlice[i]) < 2 {
			//fmt.Println("e")
			for j := i; j >= 0; j-- {
				if CompareOperators(v, OperatorsSlice[j]) < 2 && !strings.Contains(strLeftParenthesis, OperatorsSlice[j]) {
					//fmt.Println("meet", v, "output", OperatorsSlice[j])
					ret += OperatorsSlice[j] + "|"
					i--
				} else if CompareOperators(v, OperatorsSlice[j]) == 2 {
					i++
					OperatorsSlice[i] = v
					j = 0
				} else {
					i--
					j = 0
				}
			}

			if i < 0 {
				i++
				OperatorsSlice[i] = v
			}

		} else {
			//fmt.Println("f")
			i++
			OperatorsSlice[i] = v
		}
	}

	for ; i >= 0; i-- {
		ret += OperatorsSlice[i] + "|"
	}

	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func formatLength(l int) (ret string) {
	lstr := strconv.Itoa(l)

	for i := len(lstr) + 1; i <= 12; i++ {
		ret += "0"
	}

	return ret + lstr
}
