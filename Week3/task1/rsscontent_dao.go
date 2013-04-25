package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func AddRssData(cRssContent *RssContent) {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT rsscontent SET atitle=?,adesc=?,alink=?,atime=?,srcid=?,ahash1=?,ahash2=?")
	checkError(err)

	res, err := stmt.Exec(cRssContent.Atitle, cRssContent.Adesc, cRssContent.Alink, cRssContent.Atime, cRssContent.Srcid, cRssContent.Ahash1, cRssContent.Ahash2)
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id, cRssContent.Srcid, cRssContent.Atitle)

}

func GetAllRssData(startIndex int32, itemLimit int32) *[]RssContent {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	stmt, err := db.Prepare("SELECT rc.id,rc.atitle,rc.adesc,rc.alink,rc.atime,rs.rssname FROM rsscontent rc inner join rsssource rs on rc.srcid = rs.id ORDER BY rc.atime desc" + " LIMIT ?,?")
	checkError(err)

	rows, err := stmt.Query(startIndex, itemLimit)
	checkError(err)

	RssContentLs := make([]RssContent, 0, itemLimit)

	for rows.Next() {
		var id int
		var atitle string
		var adesc string
		var alink string
		var atimeStr string
		var rssname string
		_ = rows.Scan(&id, &atitle, &adesc, &alink, &atimeStr, &rssname)

		//fmt.Println("|" + atimeStr + "|")
		//fmt.Println(layoutParser("%Y-%m-%d %H:%M:%S"))
		//atime, _ := Parse("%Y-%m-%d %H:%M:%S", atimeStr)
		RssContentLs = append(RssContentLs, RssContent{Id: id, Atitle: atitle, Adesc: adesc, Alink: alink, Atime: ParseTime(atimeStr), SrcName: rssname})
	}

	fmt.Println(len(RssContentLs))

	return &RssContentLs
}

func GetAllRssDataAmount() int {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(1) FROM rsscontent")
	checkError(err)

	var recordAmount int

	for rows.Next() {

		_ = rows.Scan(&recordAmount)
	}

	return recordAmount
}

func GetRssDataAmount(ahash1 uint32, ahash2 uint32) int {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	stmt, err := db.Prepare("SELECT COUNT(1) FROM rsscontent WHERE ahash1=? and ahash2=?")
	checkError(err)

	rows, err := stmt.Query(ahash1, ahash2)
	checkError(err)

	var recordAmount int

	for rows.Next() {

		_ = rows.Scan(&recordAmount)
	}

	return recordAmount
}
