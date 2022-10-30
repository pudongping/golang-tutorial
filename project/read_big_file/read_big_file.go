// 原文地址：https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// 原作者用的日志时间格式
// const timeLayout = "2006-01-02T15:04:05.0000Z"

// 改成我自己的日志时间格式后
const timeLayout = "2006-01-02 15:04:05"

// 日期时间正则表达式，用于从每一行日志中取出时间
const timeRe = `\d{4}-\d{1,2}-\d{1,2}\s\d{1,2}:\d{1,2}:\d{1,2}`

// go run read_big_file.go -f "2022-10-11 11:00:00" -t "2022-10-11 20:01:46" -i "./2022-10-11.log"
func main() {

	s := time.Now()
	args := os.Args[1:]
	if len(args) != 6 { // for format  LogExtractor.exe -f "From Time" -t "To Time" -i "Log file directory location"
		fmt.Println("Please give proper command line arguments")
		return
	}
	startTimeArg := args[1]
	finishTimeArg := args[3]
	fileName := args[5]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return
	}

	defer file.Close() // close after checking err

	queryStartTime, err := time.Parse(timeLayout, startTimeArg)
	if err != nil {
		fmt.Println("Could not able to parse the start time", startTimeArg)
		return
	}

	queryFinishTime, err := time.Parse(timeLayout, finishTimeArg)
	if err != nil {
		fmt.Println("Could not able to parse the finish time", finishTimeArg)
		return
	}

	filestat, err := file.Stat()
	if err != nil {
		fmt.Println("Could not able to get the file stat")
		return
	}

	fileSize := filestat.Size()
	fmt.Printf("文件 %s 大小为 %v kb \n", filestat.Name(), fileSize)
	offset := fileSize - 1
	lastLineSize := 0

	for {
		b := make([]byte, 1)
		// n 为读取的字节数，比如 100
		n, err := file.ReadAt(b, offset)
		if err != nil {
			fmt.Println("Error reading file ", err)
			break
		}
		char := string(b[0])
		if char == "\n" {
			break
		}
		offset--
		lastLineSize += n
	}

	lastLine := make([]byte, lastLineSize) // 最后一行数据
	_, err = file.ReadAt(lastLine, offset+1)

	if err != nil {
		fmt.Println("Could not able to read last line with offset", offset, "and lastline size", lastLineSize)
		return
	}

	// 原作者的日志示例为
	// 2020-01-31T20:12:38.1234Z, Some Field, Other Field, And so on, Till new line,...\n
	// logSlice := strings.SplitN(string(lastLine), ",", 2) // 获取最后一行数据的数据
	// logCreationTimeString := logSlice[0]                 // 这里取的是最后一行数据的时间
	re := regexp.MustCompile(timeRe)
	logCreationTimeString := re.FindString(string(lastLine)) // 这里取的是最后一行数据的时间
	fmt.Println(logCreationTimeString)

	lastLogCreationTime, err := time.Parse(timeLayout, logCreationTimeString)
	if err != nil {
		fmt.Println("无法解析时间 ==>", err)
	}

	// 日志文件中的最后一行时间是否在所需查询开始时间之后？并且在所需查询结束时间之前？
	// 也就是说只有符合以下时间区间的，才做读取该文件，否则不读
	// queryStartTime <= lastLogCreationTime <= queryFinishTime
	isStart := lastLogCreationTime.After(queryStartTime) || lastLogCreationTime.Equal(queryStartTime)
	isEnd := lastLogCreationTime.Before(queryFinishTime) || lastLogCreationTime.Equal(queryFinishTime)
	if isStart && isEnd {
		Process(file, queryStartTime, queryFinishTime)
	}

	fmt.Printf("\n\u001B[32m程序执行总耗时： %v\u001B[0m\n", time.Since(s))
}

func Process(f *os.File, start time.Time, end time.Time) error {

	// sync.Pool是一个强大的对象池，可以重用对象来减轻垃圾收集器的压力。我们将重用各个分片的内存，以减少内存消耗，大大加快我们的工作。
	// Go Routines帮助我们同时处理缓冲区块，这大大提高了处理速度。
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 250*1024)
		return lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	// 创建一个具有默认大小缓冲的 Reader
	r := bufio.NewReader(f)

	var wg sync.WaitGroup

	for {
		buf := linesPool.Get().([]byte) // 块大小

		// 读取到达结尾时，返回值 n 将为 0 而 err 将为 io.EOF
		n, err := r.Read(buf) // 将块加载到缓冲区
		buf = buf[:n]

		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return err
		}

		// 读取直到第一次遇到 `\n` 字节，返回一个包含已读取的数据和 `\n` 字节的切片
		// 如果 ReadBytes 方法在读取到 `\n` 之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是 io.EOF）
		// 当且仅当 ReadBytes 方法返回的切片不以 `\n` 结尾时，会返回一个非 nil 的错误
		nextUntilNewline, err := r.ReadBytes('\n') // 读取整行

		if err != io.EOF {
			buf = append(buf, nextUntilNewline...)
		}

		wg.Add(1)
		go func() {
			// 同时处理每个块
			ProcessChunk(buf, &linesPool, &stringPool, start, end)
			wg.Done()
		}()

	}

	wg.Wait()
	return nil
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, start time.Time, end time.Time) {

	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs) // 放回字符串池

	chunkSize := 300 // 处理线程中的 300 条日志
	n := len(logsSlice)
	noOfThread := n / chunkSize

	// 检查是否溢出
	if n%chunkSize != 0 {
		noOfThread++
	}

	// 遍历块
	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done() // to avaoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i] // 每一行文本内容
				if len(text) == 0 {
					continue
				}

				// 原作者的日志示例为
				// 2020-01-31T20:12:38.1234Z, Some Field, Other Field, And so on, Till new line,...\n
				// logSlice := strings.SplitN(text, ",", 2)
				// logCreationTimeString := logSlice[0]
				re := regexp.MustCompile(timeRe)             // 因为是正则，因此在这里速度可能会有影响，建议考虑通过字符串截取去获取时间
				logCreationTimeString := re.FindString(text) // 这里取出每一行记录的时间

				logCreationTime, err := time.Parse(timeLayout, logCreationTimeString)
				if err != nil {
					fmt.Printf("\n Could not able to parse the time :%s for log : %v", logCreationTimeString, text)
					return
				}

				if (logCreationTime.After(start) || logCreationTime.Equal(start)) && (logCreationTime.Before(end) || logCreationTime.Equal(end)) {
					// 这里你也可以按照你的要求去过滤文本内容
					// 这里筛选出的是时间区间内的文本内容
					// fmt.Println(text)
					fmt.Printf("===>%s<===\n", text)
				}
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg2.Wait()
	logsSlice = nil
}
