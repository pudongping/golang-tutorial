# panic 和 recover

- [panic 的使用机制](./panic.go)
- [recover 的执行机制](./recover.go)
- [defer panic recover 结合正确处理错误方式](./panic_recover.go)

## recover 不是万能的

[详见](./recover1.go)

recover 并非万能，并不能够捕获到所有的错误，它只针对用户态下的 `panic` 关键字有效。在 Go 语言中，是存在着这些无法恢复的 ”恐慌“，例如像是 fatalthrow、fatalpanic 等等方法，
因此自然而然使用 recover 就无法捕获到了，因为它是直接退出程序（比如，调用了 exit() 函数），结果是中断程序。

- panic 只能触发当前 Goroutine 的 defer 调用，在 defer 调用中如果存在 recover ，那么就能够处理其所抛出的恐慌事件。但是需要注意的是在其它 Goroutine 中的 defer 是对其没有用的，并不支持跨协程（goroutine），需要分清楚。
- 想捕获/处理 panic 所造成的恐慌，recover 必须与 defer 配套使用，否则无效。
- 在 Go 语言中，是存在着无法处理的致命错误方法的，例如：fatalthrow、fatalpanic 方法，一般会在并发写入 map 等等处理时抛出，需要谨慎。