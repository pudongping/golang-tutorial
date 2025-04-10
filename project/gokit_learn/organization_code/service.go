package main

// service.go

import (
	"context"
	"errors"
)

// AddService 列出当前服务所有RPC方法的接口类型
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// addService 实现AddService接口
type addService struct {
	// ...
}

var (
	// ErrEmptyString 两个参数都是空字符串的错误
	ErrEmptyString = errors.New("两个参数都是空字符串")
)

// Sum 返回两个数的和
func (addService) Sum(_ context.Context, a, b int) (int, error) {
	// 业务逻辑
	return a + b, nil
}

// Concat 拼接两个字符串
func (addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}
	return a + b, nil
}

// NewService 创建一个add service
func NewService() AddService {
	return &addService{}
}
