package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func AddSettingData(rssname string, rssurl string) {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT rsssource SET rssname=?,rssurl=?")
	checkError(err)

	res, err := stmt.Exec(rssname, rssurl)
	checkError(err)
	fmt.Println(rssname, rssurl)
	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)

}

func GetAllSettingData() *[]RssSource {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	//stmt, _ := db.Prepare("SELETE * FROM rsssource")

	rows, err := db.Query("SELECT * FROM rsssource")
	checkError(err)

	RssSourceLs := make([]RssSource, 0, 25)

	for rows.Next() {
		var id int
		var rssname string
		var rssurl string
		_ = rows.Scan(&id, &rssname, &rssurl)

		RssSourceLs = append(RssSourceLs, RssSource{id, rssname, rssurl})
	}

	if len(RssSourceLs) == 0 {
		//RssSourceLs = append(RssSourceLs, RssSource{0, "a", "b"})
	}
	fmt.Println(len(RssSourceLs))

	return &RssSourceLs
}

func DelSettingData(id int) {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM rsssource WHERE id=?")

	stmt.Exec(id)
}
