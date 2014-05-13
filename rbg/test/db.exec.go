package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db, _ = sql.Open("mysql", "root:wenbin@/brs?charset=utf8")

func Count(i int) {

	s := "insert into "
	if i == 1 {
		s += "pa1"
	} else if i == 2 {
		s += "pa2"
	} else if i == 3 {
		s += "pa3"
	} else {
		s += "pa4"
	}

	s += " (text, time) values ('learsu', ?)"
	t1 := time.Now()
	for i := 0; i <= 3; i++ {
		now := time.Now().Unix()

		//DB insert直接插入数据库 不适合守候程序 预编译过的速度更快
		rs, err := db.Exec(s, now)
		if err == nil {
			id, _ := rs.LastInsertId()
			fmt.Println(">>", i, id)
		} else {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}
	fmt.Println(time.Now().Sub(t1))
}
func main() {
	for i := 0; i < 100; i++ {
		go Count(i)
	}
	select {}
}
