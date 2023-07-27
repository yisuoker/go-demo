package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yisuoker/go-demo/conf"
)

var db *sql.DB

func initDB() (err error) {
	if err = conf.Init(""); err != nil {
		fmt.Printf("Failed to load the configuration file, err: %v\n", err)
		return
	}
	// fmt.Printf("%#v\n", conf.Cfg)

	cfg := conf.Cfg.DbCfg
	// fmt.Printf("%#v\n", cfg)

	// dsn := "root:123456@tcp(127.0.0.1:3306)/go-demo?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.Charset,
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 单行查询
func queryRow() {
	sqlStr := "select * from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("query row failed, err:%v\n", err)
		return
	}
	fmt.Printf("query result %#v\n", u)
}

// 多行查询
func queryRows() {
	sqlStr := "select id,name from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query rows failed, err:%v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.id, &u.name); err != nil {
			fmt.Printf("scam row failed, err:%v\n", err)
		}
		fmt.Printf("query result %#v\n", u)
	}
}

// 插入数据
func insertRow() {
	sqlStr := "insert into user(name,age) values (?,?)"
	res, err := db.Exec(sqlStr, "zhangsan", 20)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}

	fmt.Println("insert success, id is ", id)
}

// 更新数据
func updateRow() {
	sqlStr := "update user set age=? where id=?"
	res, err := db.Exec(sqlStr, 30, 3)
	if err != nil {
		fmt.Printf("update row failed, err:%v\n", err)
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Println("update success, affeted row:", n)
}

// 删除数据
func deleteRow() {
	sqlStr := "delete from user where id=?"
	res, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete row failed, err:%v\n", err)
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Println("delete success, affeted row:", n)
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init db failed, err:%v", err)
		return
	}

	// err := initDB()
	// if err != nil {
	// 	fmt.Printf("init db failed, err:%v", err)
	// 	return
	// }

	defer db.Close()
	fmt.Println("connect success ...")

	queryRow()
	// queryRows()
	// insertRow()
	// updateRow()
	// deleteRow()
}
