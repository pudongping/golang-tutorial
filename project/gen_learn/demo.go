package main

import (
	"context"
	"fmt"
	"time"

	"go-tutorial/project/gen_learn/dal"
	"go-tutorial/project/gen_learn/dal/model"
	"go-tutorial/project/gen_learn/dal/query"
)

const MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug()
}

func handleError(err error) {
	if err != nil {
		panic(fmt.Sprintf("出现错误-> %+v", err))
	}
}

func main() {
	// 设置默认 DB 对象
	query.SetDefault(dal.DB)

	// 创建一个新用户
	now := time.Now().UnixNano()
	user1 := model.User{
		Name:      "Alex",
		Gender:    1,
		Version:   1,
		CreatedAt: int32(now),
		UpdatedAt: int32(now),
	}
	err := query.User.WithContext(context.Background()).Create(&user1)
	handleError(err)

	// 查询用户
	user2, err := query.User.WithContext(context.Background()).First()
	handleError(err)
	fmt.Printf("查询到的用户: %+v\n", user2)

	// 更新用户
	// 更新的时候，更新姓名和版本号加1
	ret, err := query.User.WithContext(context.Background()).
		Where(query.User.UserID.Eq(user1.UserID)).
		Updates(model.User{
			Name:    "Alex1",
			Version: user1.Version + 1,
		})
	handleError(err)
	fmt.Printf("更新的行数: %d\n", ret.RowsAffected)

	// 查看更新后的结果
	user3, err := query.Q.User.WithContext(context.Background()).
		Where(query.User.UserID.Eq(user1.UserID)).
		First()
	handleError(err)
	fmt.Printf("更新后的用户: %+v\n", user3)

	// 删除用户
	ret, err = query.User.WithContext(context.Background()).
		Where(query.User.UserID.Eq(user1.UserID)).
		Delete()
	handleError(err)
	fmt.Printf("删除的行数: %d\n", ret.RowsAffected)

	// 查看删除后的结果
	user4, err := query.Q.User.WithContext(context.Background()).
		Where(query.User.UserID.Eq(user1.UserID)).
		First()
	handleError(err)
	if user4 == nil {
		fmt.Println("用户已被删除")
	} else {
		fmt.Printf("删除后的用户: %+v\n", user4)
	}
}
