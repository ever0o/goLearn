package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:1234@(localhost:3306)/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("insert into my_job(department_id, job_name)values(?,?)")
	checkErr(err)

	res, err := stmt.Exec("4", "测试员2")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Print(id)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
