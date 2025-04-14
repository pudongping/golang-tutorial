//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"go-tutorial/project/wire_learn/wire_pkg_demo"

	"github.com/google/wire"
)

func InitializeC(ctx context.Context, name string, age int) (wire_pkg_demo.C, error) {

	// 声明一个注入器函数
	wire.Build(wire_pkg_demo.NewA, wire_pkg_demo.NewB, wire_pkg_demo.NewC)

	// 返回值无关紧要，只要类型正确即可
	return wire_pkg_demo.C{}, nil
}
