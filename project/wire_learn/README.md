# wire

wire 中有两个核心概念：提供者（provider）和注入器（injector）。

### provider

Wire 中的提供者就是一个可以产生值的普通函数。

### injector

应用程序中是用一个注入器来连接提供者，注入器就是一个按照依赖顺序调用提供者。

## 安装 wire 命令行工具

```bash
go install github.com/google/wire/cmd/wire@latest
```

在 `wire.go` 同级目录下执行 `wire` 命令，wire 会自动生成一个 `wire_gen.go` 文件，文件中包含了所有的依赖关系。

```bash
wire
```

## 运行并测试代码

```bash
go run main.go wire_gen.go
```