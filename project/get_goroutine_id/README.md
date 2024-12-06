# 获取 goroutine ID

Go 标准库中没有直接获取 goroutine ID 的方法，但是可以通过从 `runtime.Stack` 中获取 goroutine ID。

但是，这种方法并不是很好，因为 `runtime.Stack` 会导致性能下降，所以不推荐在生产环境中使用。

我们可以直接采用第三方扩展包 `github.com/petermattis/goid` 来获取 goroutine ID。

因为这个包是通过 C 和 汇编来实现的，所以性能会更好。