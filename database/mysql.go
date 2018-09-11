package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/go_test?charset=utf8")
	CheckErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	CheckErr(err)
	res, err := stmt.Exec("sunlong", "研发部门", "2018-09-11")
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)
	res, err = stmt.Exec("sunlong_update", id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)
	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	CheckErr(err)
	res, err = stmt.Exec(id)
	CheckErr(err)
	affect, err = res.RowsAffected()
	CheckErr(err)
	fmt.Println(affect)

	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 总结：
// sql.Open()函数用来打开一个注册过的数据库驱动
// db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
// db.Query()函数用来直接执行Sql返回Rows结果。
// stmt.Exec()函数用来执行stmt准备好的SQL语句