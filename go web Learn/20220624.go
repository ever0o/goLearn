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

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update my_job set department_id = ? , job_name = ? where id = ?")
	checkErr(err)

	res, err = stmt.Exec("1", "Bob", "10")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("select * from my_job")
	checkErr(err)

	for rows.Next() {
		var id int
		var department_id int
		var job_name string
		err = rows.Scan(&id, &department_id, &job_name)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(department_id)
		fmt.Println(job_name)
	}

	//删除数据
	stmt, err = db.Prepare("delete from my_job where id = ?")
	checkErr(err)

	res, err = stmt.Exec("7")
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
