package main

import (
	"bufio"
	//"bytes"
	//"encoding/binary"
	"fmt"
	"io"
	//"math"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	NotPV     []string = []string{".css", ".js", ".class", ".gif", ".jpg", ".jpeg", ".png", ".bmp", ".ico", "rss", "xml", "swf"}
	rest      string   = "deallog"
	domainMap          = make(map[string]*os.File)
	icou      int      = 0
	allcou    int      = 0
	//addr      *net.UDPAddr
	logstr chan string = make(chan string)
	db     *DB
)

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			fmt.Println(f.Name())
			return nil
		}
		if strings.Contains(f.Name(), ".log") && f.Name() != (rest+".log") {
			fmt.Println(f.Name())
			fmt.Println(f.ModTime())
			go analysis(path, &f)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	//fmt.Println(icou)
	//fmt.Println(allcou)
}

func analysis(filename string, finfo *os.FileInfo) {

	f, _ := os.Open(filename)

	defer f.Close()

	r := bufio.NewReaderSize(f, 1024*1024*10)

	for {

		strbyte, _, _ := r.ReadLine()

		if strbyte == nil {
			//fmt.Println("The End" + time.Now().String())
			break
		}

		str := string(strbyte)
		splitLog(str)

	}

	//runtime.GC()

}

func splitLog(logContent string) {

	ipStrArr := strings.Split(logContent, " ")
	time_taken, _ := strconv.Atoi(ipStrArr[len(ipStrArr)-1])

	allcou = allcou + 1

	if time_taken > 250000 {
		//fmt.Println(logContent)
		icou = icou + 1

		//var str string
		//fmt.Scanf("%s", &str)
		logstr <- logContent + "\n"

	}

	//return
}

func client() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, _ := net.DialUDP("udp", nil, addr)

	defer conn.Close()

	for {
		str := <-logstr
		//fmt.Println(str)
		io.WriteString(conn, str)
	}
}

func server() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
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

func init() {
	db, err := sql.Open("sqlite3", "./agent.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	//查询数据
	rows, err := db.Query("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?;", "dealfilelog")
	checkErr(err)

	var tabAmount int = 0
	for rows.Next() {
		err = rows.Scan(&tabAmount)
		checkErr(err)
	}
	fmt.Println(tabAmount)
	if tabAmount == 0 {
		sreCreateTable := "CREATE TABLE `dealfilelog` (" +
			"`logfileid` INTEGER PRIMARY KEY AUTOINCREMENT," +
			"`logfilename` VARCHAR NULL," +
			"`offset` INTEGER NULL," +
			"`filemodifydate` DATE NULL," +
			"`dealdate` DATE NULL" +
			");"

		stmt, err := db.Prepare(sreCreateTable)
		checkErr(err)

		_, err = stmt.Exec()
		checkErr(err)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Begin" + time.Now().String())
	//go server()
	go client()
	getFilelist("./")
	//analysis("u_ex14021620.log")
	fmt.Println("The End" + time.Now().String())
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
