# golang 学习笔记

## 基础

### 1. 数据类型

- [数据类型](./knowledge/chapter01/01-datatype)
- [变量](./knowledge/chapter01/02-variable)
- [常量和枚举 constant](./knowledge/chapter01/03-constant)
- [数据类型之间的转化](./knowledge/chapter01/04-type-conversation)
- [数组 array](./knowledge/chapter01/05-array)
- [切片 slice](./knowledge/chapter01/06-slice)
- [字典 map](./knowledge/chapter01/07-map)
- [指针 pointer](./knowledge/chapter01/08-pointer)

### 2. 流程控制

- [if 条件语句](./knowledge/chapter02/01-condition)
- [switch case 分支语句](./knowledge/chapter02/02-switch-case)
- [for 循环语句](./knowledge/chapter02/03-for-loop)
- [break 与 continue 跳转语句](./knowledge/chapter02/04-jump-statement)

### 3. 函数式编程

- [函数](./knowledge/chapter03/01-function)
- [参数传递、变长参数与多返回值](./knowledge/chapter03/02-func-params-and-return-values)
- [匿名函数与闭包](./knowledge/chapter03/03-anonymous-function-and-closure)
- [通过高阶函数实现装饰器模式](./knowledge/chapter03/04-decorator-implement-by-high-order-function) 
- [递归函数及性能调优](./knowledge/chapter03/05-recursive-function-and-optimization)
- [Map-Reduce-Filter 模式处理集合元素](./knowledge/chapter03/06-func-map-reduce-filter-mode)
- [基于管道技术实现函数的流式调用](./knowledge/chapter03/07-func-chaining-with-pipeline)

### 4. 面向对象

- [类型系统概述](./knowledge/chapter04/01-type-system)
- [类的定义、初始化和成员方法](./knowledge/chapter04/02-struct-and-class)
- [通过组合实现类的封装、继承、多态和方法重写](./knowledge/chapter04/03-oop-with-type-composite)
- [类属性和成员方法的可见性](./knowledge/chapter04/04-class-props-methods-visibility)
- [接口定义及实现](./knowledge/chapter04/05-interface)
- [接口赋值](./knowledge/chapter04/06-interface-assignment)
- [类型断言](./knowledge/chapter04/07-type-assertion)
- [空接口、反射和泛型](./knowledge/chapter04/08-empty-interface-reflection-and-generic)
- [import 导包和 init 方法调用流程](./knowledge/chapter04/09-init)

### 5. 错误处理

- [error 类型](./knowledge/chapter05/01-error)
- [defer](./knowledge/chapter05/02-defer)
- [panic 和 recover](./knowledge/chapter05/03-panic-and-recover)

### 6. 并发

- [goroutine](./knowledge/chapter06/01-goroutine)
- [channel](./knowledge/chapter06/02-channel)

## 进阶

- [Go 大杀器之性能剖析 PProf](./advanced/PProf_example)
- [Go 大杀器之跟踪剖析 trace](./advanced/trace_example)
- [用 GODEBUG 看调度跟踪 GPM](./advanced/GODEBUG_GPM_example)
- [用 GODEBUG 看 GC](./advanced/GODEBUG_GC_example)
- [Go 进程诊断工具 gops](./advanced/gops_example)
- [公开和发布度量指标](./advanced/public_publish_Metrics)
- [逃逸分析 - 变量在哪儿](./advanced/escape_analysis)

## 技巧

- [内存对齐：尽量减少内存占用](./skill/memory_alignment)
- [json 操作技巧：将数字类型传递给前端造成数字失真问题、忽略嵌套结构体空值字段](./skill/json_demo)

## 学以致用小项目和小示例

- [简单的计算器](./project/calc)
- [精简的即时通讯示例](./project/IM-System)
- [时间操作大全](./project/time_helper)
- [使用第三方包 olivere/elastic 操作 elasticsearch](./project/es_demo)
- [简单封装原生 http 客户端请求](./project/http_client)
- [高效快速读取超大日志文件](./project/read_big_file)
- [在 Go 项目中获取可靠的项目根目录](./project/get_root_path)
- [基于 redis 实现异步队列以及异步延迟队列](./project/redis_demo)
- [基于 imap 协议解析邮件内容](./project/mail_parse)
- [文件分片（可用于分片上传的前身）](./project/file_handler)
- [第三方包 rabbitmq/amqp091-go 操作 RabbitMQ](./project/rabbitmq_demo)
- [找出 Redis 中的 Big key](./project/redis_big_key)
- [抓取微信“图片/文字”类型中的图片](./project/crawler_wx_tuwen_img)
- [学习 sqlx 示例](./project/sqlx_demo)
- [使用 go-redis 来操作 redis](./project/redis_demo)
- [使用 viper 来读取配置文件](./project/viper_demo)
- [标准库 log 以及 zap 日志库使用](./project/log_demo)
- [优雅关机和平滑重启](./project/shutdown_demo)
- [分布式唯一ID——雪花算法](./project/snow_flake)
- [使用 go-playground/validator 来做参数校验](./project/validator_demo)
- [使用 golang-jwt/jwt 做鉴权](./project/jwt_demo)
- [接口速率限制：漏桶算法、令牌桶算法](./project/ratelimit_demo)
- [图片相似度对比](./project/compare_image_similarity)
- [读取/生成 二维码](./project/qr_code_demo)
- [获取 goroutine ID](./project/get_goroutine_id)

## Project supported by JetBrains

Many thanks to Jetbrains for kindly providing a license for me to work on this and other open-source projects.

<a href="https://jb.gg/OpenSourceSupport" target="_blank">
<img width="150px" src="https://s2.loli.net/2022/10/10/xlWUgcwi32J4oBT.png" alt="jb_beam">
</a>

## LICENSE

MIT