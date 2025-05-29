# redis 相关示例代码

- [redis 异步队列、延迟队列的实现](./async_queue.go)
- [遍历扫描 key](./get_keys.go)
- [redis pipeline 示例](./pipeline.go)

## 优先级队列

redis 实现优先级队列思路：

blpop 有多个键时，blpop 会从左到右依次检查每个键，直到找到一个非空的键为止。因此，可以将优先级队列的键按照优先级从高到低的顺序排列。

例如，使用 redis 命令示例
```
LPUSH queue_high "task1"
LPUSH queue_medium "task2"
LPUSH queue_low "task3"
```
这样，当使用 `BLPOP queue_high queue_medium queue_low` 时，redis 会优先返回 `queue_high` 中的任务，如果该队列为空，则检查 `queue_medium`，最后检查 `queue_low`。