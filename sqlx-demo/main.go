package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yisuoker/go-demo/conf"
)

var db *sqlx.DB

func initDB() (err error) {
	if err = conf.Init(""); err != nil {
		fmt.Printf("Failed to load the configuration file, err: %v\n", err)
		return
	}
	// fmt.Printf("%#v\n", conf.Cfg)

	cfg := conf.Cfg.DbCfg
	// fmt.Printf("%#v\n", cfg)

	// dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.Charset,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return nil
}

type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "wangwu", 28)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("connect success ...")

	// queryRowDemo()
	queryMultiRowDemo()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
}
