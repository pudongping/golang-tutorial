package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 9527, Name: "Alex", Avatar: "https://www.baidu.com"}
}

func main() {
	_ = GetUserInfo()
}
