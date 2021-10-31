# 递归函数及性能调优

- [递归的实现 - 斐波那契数列](./recursive.go)
- [通过内存缓存技术优化递归函数性能](./memoization_recursive.go)
- [通过尾递归优化递归函数性能](./tail_recursive.go)

我们知道函数调用底层是通过栈来维护的，对于递归函数而言，如果层级太深，
同时保存成百上千的调用记录，会导致这个栈越来越大，消耗大量内存空间，
严重情况下会导致栈溢出（stack overflow），为了优化这个问题，
可以引入尾递归优化技术来重用栈，降低对内存空间的消耗，提升递归函数性能。

**尾调用** 是指一个函数的最后一个动作是调用一个函数（只能是一个函数调用，不能有其他操作，比如函数相加、乘以常量等）

尾递归优化版递归函数性能要优于内存缓存技术优化版，并且不需要借助额外的内存空间保存中间结果，因此从性能角度看是更好的选择。