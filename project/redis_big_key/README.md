# 如何找到 Redis 的 Big key？

## 1. 什么是 Redis 的 Big key？

Redis 的 Big key 是指占用内存较大的 key，通常是大 List、大 Set、大 Hash、大 String 等等。

比如说：

- 字符串类型。如：超过 1 MB 的 key，就是一个 Big key。
- 非字符串类型。如：一个包含 100 万个元素的 List，占用内存 100 MB，那么这个 List 就是一个 Big key。

**具体的规定根据每个公司的实际情况而定。**

## 2. 为什么要找到 Redis 的 Big key？

- 内存空间不均匀：如果 Redis 的 Big key 占用了大量的内存，那么就会导致内存空间不均匀，从而导致 Redis 的内存不足。
- 查询时阻塞：因为 Redis 单线程特性，如果操作某个 Big key，耗时比较久，则后面的请求会被阻塞。
- 过期时阻塞：如果 Big key 设置了过期时间，当过期后，这个 key 会被删除，假如没有使用过期异步删除，就会存在阻塞 Redis 的可能性，并且慢查询中查不到（因为这个删除是内部循环事件）

## 3. 如何找到 Redis 的 Big key？

### 直接使用 `redis-cli` 命令，分析大致的情况

```bash
# 我这里在第二个数据库上做的测试，因此需要指定数据库 `-n 2`
$ redis-cli -p 6379 -n 2 --bigkeys
```

大致的结果如下：

```bash
# Scanning the entire keyspace to find biggest keys as well as
# average sizes per key type.  You can use -i 0.1 to sleep 0.1 sec
# per 100 SCAN commands (not usually needed).

[00.00%] Biggest set    found so far '"large_set_key"' with 201 members
[00.00%] Biggest list   found so far '"large_list_key"' with 201 items
[00.00%] Biggest string found so far '"large_string_key"' with 5242880 bytes
[00.00%] Biggest hash   found so far '"large_hash_key"' with 201 fields
[00.00%] Biggest zset   found so far '"large_zset_key"' with 201 members

-------- summary -------

# 一共扫描了 5 个 key
Sampled 5 keys in the keyspace!
# 所有 key 的总长度是 71 字节，平均长度为 14.20 字节
Total key length in bytes is 71 (avg len 14.20)

Biggest   list found '"large_list_key"' has 201 items
Biggest   hash found '"large_hash_key"' has 201 fields
Biggest string found '"large_string_key"' has 5242880 bytes
Biggest    set found '"large_set_key"' has 201 members
Biggest   zset found '"large_zset_key"' has 201 members

# 每一种 key 情况的总览，某种类型的 key 占用内存的百分比，平均大小
1 lists with 201 items (20.00% of keys, avg size 201.00)
1 hashs with 201 fields (20.00% of keys, avg size 201.00)
1 strings with 5242880 bytes (20.00% of keys, avg size 5242880.00)
0 streams with 0 entries (00.00% of keys, avg size 0.00)
1 sets with 201 members (20.00% of keys, avg size 201.00)
1 zsets with 201 members (20.00% of keys, avg size 201.00)
```

![redis-cli 扫描 big key](https://upload-images.jianshu.io/upload_images/14623749-a06d75c4c7c12cef.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后可以使用 `memory usage` 命令查看具体的内存占用情况：

```bash
localhost:2>MEMORY USAGE large_string_key
"6291520"
localhost:2>MEMORY USAGE large_list_key
"2240"
localhost:2>MEMORY USAGE large_set_key
"11264"
localhost:2>MEMORY USAGE large_hash_key
"4269"
localhost:2>MEMORY USAGE large_zset_key
"18968"
localhost:2>
```

![image.png](https://upload-images.jianshu.io/upload_images/14623749-fd87710c7404b132.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 通过代码脚本找到具体的 Big key

详见代码。

## 4. 如何优化 Redis 的 Big key？

### 数据结构优化

- 拆分数据：将大 key 拆分为更小的键，这可以通过拆分数据结构或者对数据进行分片来实现。
- 选择合适的数据结构：使用更适合你数据和使用场景的数据结构，比如将列表（list）转换为集合（set）、哈希（hash）或有序集合（sorted set）等。

### 数据清理

- 定期清理过期数据：确保过期数据及时清理，避免无用数据占用空间。
- 删除不必要的数据：定期清理不再需要的数据，确保 Redis 中保留的数据时真正有用的。但是要注意的是：如果直接 del，可能会导致阻塞 Redis 服务。大致有以下处理方式：
  - 使用异步删除：使用 `unlink` 异步删除，可以避免阻塞 Redis 服务，但是会导致内存占用变大。
  - 使用分批删除：将大量的删除操作分批进行，每次删除一部分，直到删除完毕。

### 内存优化

- 内存淘汰策略：调整 Redis 的内存淘汰策略，比如设置 LRU（最近最少使用）策略来淘汰不常用的键。
- 内存优化配置：调整 Redis 的内存配置参数，比如适当调整 `maxmemory` 参数，避免内存超限问题。