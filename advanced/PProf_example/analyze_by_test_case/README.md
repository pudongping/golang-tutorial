# 通过测试用例做剖析

```bash
cd analyze_by_test_case

# 采集数据，对 CPU 进行剖析
go test -bench=. -cpuprofile=cpu.profile
# 生成 png 图片进行查看
go tool pprof -png cpu.profile

# 采集数据，对内存进行剖析
go test -bench=. -memprofile=mem.profile
# 生成 png 图片进行查看
go tool pprof -png mem.profile

```