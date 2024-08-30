# 优雅关机和平滑重启

## 优雅关机

什么是优雅的关机？

优雅的关机是指在关闭服务之前，先让服务处理完当前正在处理的请求，然后再关闭服务。这样可以保证服务不会丢失请求，也不会影响到正在处理的请求。
而执行 `Ctrl + C` 或者 `kill -2 pid` 命令关闭服务，是不会等待服务处理完请求的，这样就会导致服务丢失请求。

如何实现优雅的关机？

Go 1.8 版本之后，http.Server 内置的 Shutdown() 方法就支持优雅地关机。

[代码示例](./gin_shutdown.go)

## 平滑重启

[代码示例](./gin_graceful_restart.go)