package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db, _ = sql.Open("mysql", "root:wenbin@tcp(127.0.0.1:3306)/brs?charset=utf8&timeout=30s")

func Count(i int) {

	defer db.Close()

	tname := ""
	if i == 1 {
		tname += "pa1"
	} else if i == 2 {
		tname += "pa2"
	} else if i == 3 {
		tname += "pa3"
	} else {
		tname += "pa4"
	}
	//Prepare insert入库 如果没有数据会执行Ping()保持链接或者timeout参数 prepare能保持多久未测试
	stmt, _ := db.Prepare("insert into " + tname + " (text, time) values ('learsu', ?)")
	t1 := time.Now()
	for i := 0; i <= 500; i++ {
		now := time.Now().Unix()

		res, err := stmt.Exec(now)
		if err != nil {
			fmt.Println("No user with that ID.")
			fmt.Println(err.Error())
		} else {
			id, _ := res.LastInsertId()
			fmt.Println(">>", i, id)
		}
	}
	fmt.Println(time.Now().Sub(t1))
	stmt.Close()
}

func main() {
	for i := 0; i < 100; i++ {
		go Count(i)
	}
	select {}
}
