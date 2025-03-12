# ETCD

- [ETCD 官方文档](https://etcd.io/docs/v3.5/)
- [GitHub](https://github.com/etcd-io/etcd)

## 相关资料

- [etcd：从应用场景到实现原理的全方位解读](https://www.infoq.cn/article/etcd-interpretation-application-scenario-implement-principle/)

## 安装

可以直接通过[下载二进制文件](https://etcd.io/docs/v3.5/install/)的方式来安装，更加方便。 [下载页面](https://github.com/etcd-io/etcd/releases/) 当然也可以使用 docker 来安装。

## 查看是否安装成功

```bash
$ etcd --version
etcd Version: 3.5.19
```

## 启动 etcd

```bash
# 默认会在端口 2379 上启动对客户端通信，在端口 2380 上启动对集群通信
$ etcd
```

## 测试

```bash
# 设置键值对
$ etcdctl put key1 value1
# 获取键值对
$ etcdctl get key1
# 删除键值对
$ etcdctl del key1
```

## 功能测试

### 设置键值和获取键值

```bash
$ go run get_put.go
```

### 分布式锁

```bash
# 示例来自于官方文档示例 https://pkg.go.dev/go.etcd.io/etcd/clientv3/concurrency
go run concurrency_lock.go

# 输出内容：
# acquired lock for s1
# released lock for s1
# acquired lock for s2
```