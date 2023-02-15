package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TestUser struct {
	Id       int    `db:"id"`
	Username string `db:"username""`
	Age      int    `db:"age""`
	Address  string `db:"address""`
	Position string `db:"position""`
}

var Db *sqlx.DB

/**
 * 使用 TestUser验证sqlx使用
 * test_user 表定义参考 TestUser模型定义
 *
 */
func init() {

	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:30306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func SelectById(id string) (user TestUser, error error) {
	var users []TestUser
	err := Db.Select(&users, "select * from test_user where id = ?", id)
	if err != nil {
		fmt.Println("exec failed,", err)
		error = err
		return
	}
	if len(users) > 0 {
		user = users[0]
	}
	return
}
