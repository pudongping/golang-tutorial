package main

import (
	"fmt"
	"time"
)

func main() {
	// getTime()
	// transTime()
	// calcTime()
	judgeTime()
}

func getTime() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", now, now)

	// 获取当前时间的年、月、日、时、分、秒、纳秒、微妙、毫秒
	year, month, day := time.Now().Date()
	fmt.Printf("当前时间年月日 ====> [%v][%v][%v] typeof ===> [%T][%T][%T] \n", year, month, day, year, month, day)

	nowYear := time.Now().Year()
	fmt.Printf("当前时间年 ====> %v typeof ===> %T \n", nowYear, nowYear)

	nowMonth := time.Now().Month()
	fmt.Printf("当前时间月 ====> %v typeof ===> %T \n", nowMonth, nowMonth)

	nowDay := time.Now().Day()
	fmt.Printf("当前时间日 ====> %v typeof ===> %T \n", nowDay, nowDay)

	hour, minute, second := time.Now().Clock()
	fmt.Printf("当前时间时分秒 ====> [%v][%v][%v] typeof ===> [%T][%T][%T] \n", hour, minute, second, hour, minute, second)

	nowHour := time.Now().Hour()
	fmt.Printf("当前时间时 ====> %v typeof ===> %T \n", nowHour, nowHour)

	nowMinute := time.Now().Minute()
	fmt.Printf("当前时间分 ====> %v typeof ===> %T \n", nowMinute, nowMinute)

	nowSecond := time.Now().Second()
	fmt.Printf("当前时间秒 ====> %v typeof ===> %T \n", nowSecond, nowSecond)

	nowMillisecond := time.Nanosecond.Milliseconds()
	fmt.Printf("当前时间毫秒 ====> %v typeof ===> %T \n", nowMillisecond, nowMillisecond)

	nowNanosecond := time.Now().Nanosecond()
	fmt.Printf("当前时间纳秒 ====> %v typeof ===> %T \n", nowNanosecond, nowNanosecond)

	nowMicrosecond := time.Millisecond.Microseconds()
	fmt.Printf("当前时间微秒 ====> %v typeof ===> %T \n", nowMicrosecond, nowMicrosecond)

	// 获取当前时间戳
	nowUnix := time.Now().Unix()
	fmt.Printf("当前时间时间戳（秒级别） ====> %v typeof ===> %T \n", nowUnix, nowUnix)
	nowUnixNano := time.Now().UnixNano()
	fmt.Printf("当前时间时间戳（纳秒级别） ====> %v typeof ===> %T \n", nowUnixNano, nowUnixNano)

	weekDay := time.Now().Weekday()
	fmt.Printf("当前星期几 ====> %v typeof ===> %T \n", weekDay, weekDay)
	yearDay := time.Now().YearDay()
	fmt.Printf("当前是一年中对应的第几天 ====> %v typeof ===> %T \n", yearDay, yearDay)
	location := time.Now().Location()
	fmt.Printf("当前用的时区为 ====> %v typeof ===> %T \n", location, location)

}

func transTime() {
	// 格式化时间
	ymdhis := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis, ymdhis)
	ymdhis1 := time.Now().Format("2006-01-02")
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis1, ymdhis1)
	ymdhis2 := time.Now().Format("20060102")
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis2, ymdhis2)
	ymdhis3 := time.Now().Format("15:04:05")
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis3, ymdhis3)
	ymdhis4 := time.Now().Format("150405")
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis4, ymdhis4)

	y := time.Now().Format("2006")
	fmt.Printf("当前时间年 ====> %v typeof ===> %T \n", y, y)

	m := time.Now().Format("01")
	fmt.Printf("当前时间月 ====> %v typeof ===> %T \n", m, m)

	d := time.Now().Format("02")
	fmt.Printf("当前时间日 ====> %v typeof ===> %T \n", d, d)

	h := time.Now().Format("15")
	fmt.Printf("当前时间时 ====> %v typeof ===> %T \n", h, h)

	i := time.Now().Format("04")
	fmt.Printf("当前时间分 ====> %v typeof ===> %T \n", i, i)

	s := time.Now().Format("05")
	fmt.Printf("当前时间秒 ====> %v typeof ===> %T \n", s, s)

	var timeUnix int64 = 1666599090
	goTimeUnix := time.Unix(timeUnix, 0)
	fmt.Printf("已知时间戳转 go 格式时间 ====> %v typeof ===> %T \n", goTimeUnix, goTimeUnix)
	goTimeUnixFormat := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	fmt.Printf("已知时间戳转 ymdhis 格式时间 ====> %v typeof ===> %T \n", goTimeUnixFormat, goTimeUnixFormat)

	// 获取指定时间的时间戳
	dateUnix := time.Date(2022, 10, 24, 16, 11, 30, 0, time.Local).Unix()
	fmt.Printf("2022-10-24 16:11:30 的时间戳为 ====> %v typeof ===> %T \n", dateUnix, dateUnix)

}

func calcTime() {
	// 获取当天 0 时 0 分 0 秒的时间戳
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	fmt.Printf("当天 0 时 0 分 0 秒的时间戳 ====> %v typeof ===> %T \n", startTime, startTime)
	fmt.Printf("当天 0 时 0 分 0 秒的时间 ====> %v \n", startTime.Format("2006-01-02 15:04:05"))

	// 获取当天 23 时 59 分 59 秒的时间戳
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	fmt.Printf("当天 23 时 59 分 59 秒的时间戳 ====> %v typeof ===> %T \n", endTime, endTime)
	fmt.Printf("当天 23 时 59 分 59 秒的时间 ====> %v \n", endTime.Format("2006-01-02 15:04:05"))

	currentYmdHis := currentTime.Format("2006-01-02 15:04:05")
	// 获取 1 秒钟前的时间
	t1, _ := time.ParseDuration("-1s")
	r1 := currentTime.Add(t1).Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间 %v ===> 1 秒钟前的时间 ===> %v \n", currentYmdHis, r1)

	t2, _ := time.ParseDuration("2h")
	r2 := currentTime.Add(t2).Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间 %v ===> 2 小时后的时间 ===> %v \n", currentYmdHis, r2)

	t3, _ := time.ParseDuration("1h2m30s")
	r3 := currentTime.Add(t3).Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间 %v ===> 1 小时 2 分 30 秒后的时间 ===> %v \n", currentYmdHis, r3)

	// 计算两个时间相差多少
	t4, _ := time.ParseDuration("2h")
	r4 := currentTime.Add(t4)
	t5, _ := time.ParseDuration("-1h30m")
	r5 := currentTime.Add(t5)
	fmt.Printf("相差 %v 秒 \n", r4.Sub(r5).Seconds())
	fmt.Printf("相差 %v 分钟 \n", r4.Sub(r5).Minutes())
	fmt.Printf("相差 %v 小时 \n", r4.Sub(r5).Hours())
	fmt.Printf("相差 %v 天 \n", r4.Sub(r5).Hours()/24)

}

func judgeTime() {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2022-10-24 18:18:00")
	isAfter := startTime.After(time.Now())
	isBefore := startTime.Before(time.Now())
	fmt.Printf("2022-10-24 18:18:00 是否在当前时间之前？ %v 是否在当前时间之后？ %v \n", isBefore, isAfter)

	sTime := time.Now()
	time.Sleep(time.Second * 3)
	fmt.Printf("程序开始执行时间为：%v 结束执行时间为：%v 执行了多长时间：%v 秒钟 \n", sTime.Unix(), time.Now().Unix(), time.Since(sTime))
}
