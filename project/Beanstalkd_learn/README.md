# Beanstalkd

##  docker 安装 beanstalkd

```bash
# 启动 beanstalkd 容器，默认端口为 11300
# 没有开启持久化，重启后数据会丢失，适合开发环境
docker run -d --name alex-dq \
-p 11300:11300 \
schickling/beanstalkd


 # 如果需要开启持久化
docker run -d --name alex-dq \
-p 11300:11300 \
-v $PWD/data:/data \
schickling/beanstalkd \
-b /data -f 100
```

### 为什么是持久化的？

1.  **`-b /data`**：告诉 beanstalkd 启用 binlog 机制，并将数据文件（binlog）写入容器内的 `/data` 目录。如果不加这个参数，beanstalkd 默认是在内存中运行的，重启后数据会丢失。
2.  **`-v $PWD/data:/data`**：将宿主机（你的电脑）当前目录下的 `data` 文件夹挂载到容器内的 `/data` 目录。

**只要这两个参数同时存在，容器内生成的数据文件就会实时同步保存到你宿主机的 `$PWD/data` 目录下。** 即使你删除了容器 (`docker rm`)，只要不删宿主机的 `data` 目录，下次重新启动容器挂载同一个目录，数据依然存在。

### 关于数据安全性

虽然开启了 `-b`，但 beanstalkd 默认并不是每写入一条数据就立即刷盘（fsync），而是有一定的策略（默认是根据系统调度）。如果想要更高的数据安全性（牺牲一点性能），可以添加 `-f` 参数：

- **`-f MS`**：每隔 MS 毫秒强制刷盘一次。

例如，每 100 毫秒刷盘一次：

```bash
docker run -d --name alex-dq \
-p 11300:11300 \
-v $PWD/data:/data \
schickling/beanstalkd \
-b /data -f 100
```

可以直接进入 `alex-dq` 容器执行 `beanstalkd` 命令

```bash
# 进入容器
docker exec -it alex-dq bash

beanstalkd -h

# 会输出如下内容
Use: beanstalkd [OPTIONS]

Options:
 -b DIR   wal directory --> wal 文件所在目录（默认是 /data，开启持久化时需要指定）
 -f MS    fsync at most once every MS milliseconds (use -f0 for "always fsync") --> 每隔 MS 毫秒强制刷盘一次（默认是 0，即不强制）
 -F       never fsync (default) --> 不强制刷盘（默认是开启的）
 -l ADDR  listen on address (default is 0.0.0.0) --> 监听的 IP 地址（默认是 0.0.0.0，即监听所有地址）
 -p PORT  listen on port (default is 11300) --> 监听的端口号（默认是 11300）
 -u USER  become user and group --> 切换到指定用户和用户组
 -z BYTES set the maximum job size in bytes (default is 65535) --> 最大任务大小（默认是 65535 字节）
 -s BYTES set the size of each wal file (default is 10485760) --> 每个 wal 文件的大小（默认是 10485760 字节）
            (will be rounded up to a multiple of 512 bytes) --> 会被四舍五入到最近的 512 字节的倍数 
 -c       compact the binlog (default) --> 开启 binlog 压缩（默认是开启的）
 -n       do not compact the binlog --> 不开启 binlog 压缩
 -v       show version information --> 显示版本信息
 -V       increase verbosity --> 增加日志 verbosity（默认是 0）
 -h       show this help --> 显示帮助信息
```

### 检查是否安装成功

```bash
telnet 127.0.0.1 11300

# 输入 stats 命令，如果有大量统计信息返回，则表示成功
stats

# 如果不使用 telnet 也可以直接通过查看 docker 容器日志来检查是否安装成功
docker logs alex-dq
```

![](./imgs/stats.png)

## 安装 beanstalk console Web 管理工具

```bash
# 其中 BEANSTALK_SERVERS 为 beanstalkd 的地址和端口
docker run -d \
--name alex-dq-console \
-p 2080:2080 \
-e BEANSTALK_SERVERS=192.168.1.224:11300 \
schickling/beanstalkd-console
```

可以直接通过浏览器访问 `http://localhost:2080/` 来查看 beanstalkd 的状态和队列信息。

![](./imgs/beanstalkd-console-web.png)