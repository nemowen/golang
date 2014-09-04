package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=wenbin host=192.168.0.102 "+
		"port=5432 dbname=postgres sslmode=disable")
	checkErr(err)

	//插入数据
	//_, err := db.Exec("INSERT INTO test VALUES('astaxie','研发部门')")
	//checkErr(err)
	var querystr = "AC3934812342"

	//查询数据
	rows, err := db.Query("SELECT b.number_str, b.name, similarity($1,b.number_str) a "+
		"FROM test b where similarity($2,b.number_str) > 0.1 order by a desc", querystr, querystr)
	checkErr(err)

	for rows.Next() {
		var number string
		var name string
		var va float32
		err = rows.Scan(&number, &name, &va)
		checkErr(err)
		fmt.Println(number, name, va)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
