package sqlx_demo

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 一定不能忘记导入数据库驱动
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用 MustConnect 连接不成功就直接 panic
	// db = sqlx.MustConnect("mysql", dsn)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10) // 设置数据库连接池的最大空闲连接数
	return
}

func CreateTable(db *sqlx.DB) (err error) {
	// 写SQL语句
	sqlStr := `create table if not exists users (
		id bigint primary key auto_increment,
		name varchar(20),
		age int default 1
	);`
	_, err = db.Exec(sqlStr)

	return err
}

// 插入用户并获取 ID
func insertUser(db *sqlx.DB, name string, age int) (int64, error) {
	result, err := db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", name, age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 查询单条用户记录
func getUser(db *sqlx.DB, id int64) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 查询所有用户记录
func getAllUsers(db *sqlx.DB, id int64) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users where id > ?", id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 更新用户信息
func updateUser(db *sqlx.DB, id int64, name string, age int) (int64, error) {
	result, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", name, age, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// 删除用户记录
func deleteUser(db *sqlx.DB, id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// 使用命名参数插入用户
func insertUserNamed(db *sqlx.DB, name string, age int) (int64, error) {
	query := `INSERT INTO users(name, age) VALUES(:name, :age)`
	result, err := db.NamedExec(query, map[string]interface{}{
		"name": name,
		"age":  age,
	})
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 使用命名参数查询用户
func getUsersNamed(db *sqlx.DB, name string) ([]User, error) {
	query := `SELECT * FROM users WHERE name = :name`
	var users []User
	rows, err := db.NamedQuery(query, map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.StructScan(&user)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func RunCRUD() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面

	// 创建表
	err = CreateTable(db)
	if err != nil {
		fmt.Printf("create table failed, err:%v\n", err)
		return
	}

	// 插入数据
	id, err := insertUser(db, "Alex", 18)
	if err != nil {
		fmt.Printf("insert user failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success, the id is:", id)

	// 查询单条数据
	user, err := getUser(db, id)
	if err != nil {
		fmt.Printf("get user failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", user)

	// 查询多条数据
	users, err := getAllUsers(db, 0)
	if err != nil {
		fmt.Printf("get all users failed, err:%v\n", err)
		return
	}

	fmt.Printf("users:%#v\n", users)

	// 更新数据
	rowsAffected, err := updateUser(db, id, "Alex", 20)
	if err != nil {
		fmt.Printf("update user failed, err:%v\n", err)
		return
	}

	fmt.Println("update success, affected rows:", rowsAffected)

	// 删除数据
	rowsAffected, err = deleteUser(db, id)
	if err != nil {
		fmt.Printf("delete user failed, err:%v\n", err)
		return
	}

	fmt.Println("delete success, affected rows:", rowsAffected)

	// 使用命名参数插入数据
	id, err = insertUserNamed(db, "Alex", 19)
	if err != nil {
		fmt.Printf("insert user named failed, err:%v\n", err)
		return
	}

	fmt.Println("insert named success, the id is:", id)

	// 使用命名参数查询数据
	users, err = getUsersNamed(db, "Alex")
	if err != nil {
		fmt.Printf("get users named failed, err:%v\n", err)
		return
	}

	fmt.Printf("users named:%#v\n", users)

	fmt.Println("exec SQL success")
}
