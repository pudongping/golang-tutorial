package main

import "time"

const (
	// BeanstalkdAddr Beanstalkd 服务地址
	BeanstalkdAddr = "127.0.0.1:11300"
	// TubeName 管道名称
	TubeName = "test_tube"
	// ReserveTimeout 消费者预留任务超时时间
	ReserveTimeout = 5 * time.Second
)
