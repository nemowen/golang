package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	//db, err := sql.Open("mysql", "root:wenbin@tcp(127.0.0.1:3306)/brs?charset=utf8")
	db, err := sql.Open("mysql", "root:wenbin@/brs?charset=utf8")
	CheckErr(err)
	defer db.Close()
	rows, err := db.Query("select CT3001,CT3002 from t_user")
	CheckErr(err)
	defer rows.Close()
	var CT3001 string
	var CT3002 string
	for rows.Next() {
		err = rows.Scan(&CT3001, &CT3002)
		CheckErr(err)
		fmt.Println(CT3001, CT3002)
	}
	//str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")

	str_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 14:03:33.868")

	fmt.Println(str_time)

	maps := make(map[string]string)
	maps["1"] = "v1"
	maps["2"] = "v2"
	fmt.Println(maps, len(maps))

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
