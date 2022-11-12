package main

import (
	"fmt"
	"time"
)

func main() {
	getTime()
	transTime()
	calcTime()
	judgeTime()
	adjustTimeZone()
}

func getTime() {
	// 获取当前时间
	now := time.Now()
	// 当前时间 ====> 2022-10-24 23:36:33.47472 +0800 CST m=+0.000077400 typeof ===> time.Time
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", now, now)

	// 获取当前时间的年、月、日、时、分、秒、纳秒、微妙、毫秒
	year, month, day := time.Now().Date()
	// 当前时间年月日 ====> [2022][October][24] typeof ===> [int][time.Month][int]
	fmt.Printf("当前时间年月日 ====> [%v][%v][%v] typeof ===> [%T][%T][%T] \n", year, month, day, year, month, day)

	nowYear := time.Now().Year()
	// 当前时间年 ====> 2022 typeof ===> int
	fmt.Printf("当前时间年 ====> %v typeof ===> %T \n", nowYear, nowYear)

	nowMonth := time.Now().Month()
	// 当前时间月 ====> October typeof ===> time.Month
	fmt.Printf("当前时间月 ====> %v typeof ===> %T \n", nowMonth, nowMonth)

	nowDay := time.Now().Day()
	// 当前时间日 ====> 24 typeof ===> int
	fmt.Printf("当前时间日 ====> %v typeof ===> %T \n", nowDay, nowDay)

	hour, minute, second := time.Now().Clock()
	// 当前时间时分秒 ====> [23][36][33] typeof ===> [int][int][int]
	fmt.Printf("当前时间时分秒 ====> [%v][%v][%v] typeof ===> [%T][%T][%T] \n", hour, minute, second, hour, minute, second)

	nowHour := time.Now().Hour()
	// 当前时间时 ====> 23 typeof ===> int
	fmt.Printf("当前时间时 ====> %v typeof ===> %T \n", nowHour, nowHour)

	nowMinute := time.Now().Minute()
	// 当前时间分 ====> 36 typeof ===> int
	fmt.Printf("当前时间分 ====> %v typeof ===> %T \n", nowMinute, nowMinute)

	nowSecond := time.Now().Second()
	// 当前时间秒 ====> 33 typeof ===> int
	fmt.Printf("当前时间秒 ====> %v typeof ===> %T \n", nowSecond, nowSecond)

	// 1秒(s) ＝1000毫秒(ms)
	// 1毫秒(ms)＝1000微秒 (us) ==> Milliseconds ==> 毫秒
	// 1微秒(us)＝1000纳秒 (ns)  ==> Microseconds  ==> 微秒
	// 1纳秒(ns)＝1000皮秒 (ps)  ==> Nanoseconds  ==> 纳秒
	nowNanosecond := time.Now().Nanosecond()
	// 当前时间纳秒 ====> 474905000 typeof ===> int
	fmt.Printf("当前时间纳秒 ====> %v typeof ===> %T \n", nowNanosecond, nowNanosecond)

	// 获取当前时间戳
	nowUnix := time.Now().Unix()
	// 当前时间时间戳（秒级别） ====> 1666625793 typeof ===> int64
	fmt.Printf("当前时间时间戳（秒级别） ====> %v typeof ===> %T \n", nowUnix, nowUnix)
	nowUnixNano := time.Now().UnixNano()
	// 当前时间时间戳（纳秒级别） ====> 1666625793474909000 typeof ===> int64
	fmt.Printf("当前时间时间戳（纳秒级别） ====> %v typeof ===> %T \n", nowUnixNano, nowUnixNano)

	weekDay := time.Now().Weekday()
	// 当前星期几 ====> Monday typeof ===> time.Weekday
	fmt.Printf("当前星期几 ====> %v typeof ===> %T \n", weekDay, weekDay)
	yearDay := time.Now().YearDay()
	// 当前是一年中对应的第几天 ====> 297 typeof ===> int
	fmt.Printf("当前是一年中对应的第几天 ====> %v typeof ===> %T \n", yearDay, yearDay)
	location := time.Now().Location()
	// 当前用的时区为 ====> Local typeof ===> *time.Location
	fmt.Printf("当前用的时区为 ====> %v typeof ===> %T \n", location, location)

}

func transTime() {
	// 格式化时间
	ymdhis := time.Now().Format("2006-01-02 15:04:05")
	// 当前时间 ====> 2022-10-24 23:40:31 typeof ===> string
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis, ymdhis)
	ymdhis1 := time.Now().Format("2006-01-02")
	// 当前时间 ====> 2022-10-24 typeof ===> string
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis1, ymdhis1)
	ymdhis2 := time.Now().Format("20060102")
	// 当前时间 ====> 20221024 typeof ===> string
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis2, ymdhis2)
	ymdhis3 := time.Now().Format("15:04:05")
	// 当前时间 ====> 23:40:31 typeof ===> string
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis3, ymdhis3)
	ymdhis4 := time.Now().Format("150405")
	// 当前时间 ====> 234031 typeof ===> string
	fmt.Printf("当前时间 ====> %v typeof ===> %T \n", ymdhis4, ymdhis4)

	y := time.Now().Format("2006")
	// 当前时间年 ====> 2022 typeof ===> string
	fmt.Printf("当前时间年 ====> %v typeof ===> %T \n", y, y)

	m := time.Now().Format("01")
	// 当前时间月 ====> 10 typeof ===> string
	fmt.Printf("当前时间月 ====> %v typeof ===> %T \n", m, m)

	d := time.Now().Format("02")
	// 当前时间日 ====> 24 typeof ===> string
	fmt.Printf("当前时间日 ====> %v typeof ===> %T \n", d, d)

	h := time.Now().Format("15")
	// 当前时间时 ====> 23 typeof ===> string
	fmt.Printf("当前时间时 ====> %v typeof ===> %T \n", h, h)

	i := time.Now().Format("04")
	// 当前时间分 ====> 40 typeof ===> string
	fmt.Printf("当前时间分 ====> %v typeof ===> %T \n", i, i)

	s := time.Now().Format("05")
	// 当前时间秒 ====> 31 typeof ===> string
	fmt.Printf("当前时间秒 ====> %v typeof ===> %T \n", s, s)

	// 时间戳转时间格式
	var timeUnix int64 = 1666599090
	goTimeUnix := time.Unix(timeUnix, 0)
	// 已知时间戳转 go 格式时间 ====> 2022-10-24 16:11:30 +0800 CST typeof ===> time.Time
	fmt.Printf("已知时间戳转 go 格式时间 ====> %v typeof ===> %T \n", goTimeUnix, goTimeUnix)
	goTimeUnixFormat := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	// 已知时间戳转 ymdhis 格式时间 ====> 2022-10-24 16:11:30 typeof ===> string
	fmt.Printf("已知时间戳转 ymdhis 格式时间 ====> %v typeof ===> %T \n", goTimeUnixFormat, goTimeUnixFormat)

	// 获取指定时间的时间戳
	dateUnix := time.Date(2022, 10, 24, 16, 11, 30, 0, time.Local).Unix()
	// 2022-10-24 16:11:30 的时间戳为 ====> 1666599090 typeof ===> int64
	fmt.Printf("2022-10-24 16:11:30 的时间戳为 ====> %v typeof ===> %T \n", dateUnix, dateUnix)

}

func calcTime() {
	// 获取当天 0 时 0 分 0 秒的时间戳
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	// 当天 0 时 0 分 0 秒的时间戳 ====> 2022-10-24 00:00:00 +0800 CST typeof ===> time.Time
	fmt.Printf("当天 0 时 0 分 0 秒的时间戳 ====> %v typeof ===> %T \n", startTime, startTime)
	// 当天 0 时 0 分 0 秒的时间 ====> 2022-10-24 00:00:00
	fmt.Printf("当天 0 时 0 分 0 秒的时间 ====> %v \n", startTime.Format("2006-01-02 15:04:05"))

	// 获取当天 23 时 59 分 59 秒的时间戳
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	// 当天 23 时 59 分 59 秒的时间戳 ====> 2022-10-24 23:59:59 +0800 CST typeof ===> time.Time
	fmt.Printf("当天 23 时 59 分 59 秒的时间戳 ====> %v typeof ===> %T \n", endTime, endTime)
	// 当天 23 时 59 分 59 秒的时间 ====> 2022-10-24 23:59:59
	fmt.Printf("当天 23 时 59 分 59 秒的时间 ====> %v \n", endTime.Format("2006-01-02 15:04:05"))

	// 不超过 24 小时的时间计算
	currentYmdHis := currentTime.Format("2006-01-02 15:04:05")
	// 获取 1 秒钟前的时间
	t1, _ := time.ParseDuration("-1s")
	r1 := currentTime.Add(t1).Format("2006-01-02 15:04:05")
	// 当前时间 2022-10-24 23:44:12 ===> 1 秒钟前的时间 ===> 2022-10-24 23:44:11
	fmt.Printf("当前时间 %v ===> 1 秒钟前的时间 ===> %v \n", currentYmdHis, r1)

	t2, _ := time.ParseDuration("2h")
	r2 := currentTime.Add(t2).Format("2006-01-02 15:04:05")
	// 当前时间 2022-10-24 23:44:12 ===> 2 小时后的时间 ===> 2022-10-25 01:44:12
	fmt.Printf("当前时间 %v ===> 2 小时后的时间 ===> %v \n", currentYmdHis, r2)

	t3, _ := time.ParseDuration("1h2m30s")
	r3 := currentTime.Add(t3).Format("2006-01-02 15:04:05")
	// 当前时间 2022-10-24 23:44:12 ===> 1 小时 2 分 30 秒后的时间 ===> 2022-10-25 00:46:42
	fmt.Printf("当前时间 %v ===> 1 小时 2 分 30 秒后的时间 ===> %v \n", currentYmdHis, r3)

	// 计算两个时间相差多少
	t4, _ := time.ParseDuration("1h")
	r4 := currentTime.Add(t4 * 2) // 注意：这里是 2 小时后的时间
	t5, _ := time.ParseDuration("-1h30m")
	r5 := currentTime.Add(t5)
	// 相差 12600 秒
	fmt.Printf("相差 %v 秒 \n", r4.Sub(r5).Seconds())
	// 相差 210 分钟
	fmt.Printf("相差 %v 分钟 \n", r4.Sub(r5).Minutes())
	// 相差 3.5 小时
	fmt.Printf("相差 %v 小时 \n", r4.Sub(r5).Hours())
	// 相差 0.14583333333333334 天
	fmt.Printf("相差 %v 天 \n", r4.Sub(r5).Hours()/24)

	// 超过 24 小时之外的时间计算
	t6 := currentTime.AddDate(0, 2, 1)
	r6 := t6.Format("2006-01-02 15:04:05")
	// 当前时间 2022-10-24 23:44:12 ===> 2 个月 1 天后的时间 ===> 2022-12-25 23:44:12
	fmt.Printf("当前时间 %v ===> 2 个月 1 天后的时间 ===> %v \n", currentYmdHis, r6)

	t7 := currentTime.AddDate(0, 0, -5)
	r7 := t7.Format("2006-01-02 15:04:05")
	// 当前时间 2022-10-24 23:44:12 ===> 5 天前的时间 ===> 2022-10-19 23:44:12
	fmt.Printf("当前时间 %v ===> 5 天前的时间 ===> %v \n", currentYmdHis, r7)

}

func judgeTime() {
	now := time.Now()
	s := now.Format("2006-01-02 15:04:05")
	// 当前时间为 2022-11-12 19:45:36
	fmt.Printf("当前时间为 %s \n", s)

	// 时间比较
	// 尽可能的使用 time.ParseInLocation() 方法，因为可以手动指定时区
	// 而少用 time.Parse() 方法，因为默认的时区为 UTC（零时区）
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2022-10-24 18:18:00")
	isBefore := startTime.Before(now)
	isAfter := startTime.After(now)
	isEqual := startTime.Equal(now)
	// 2022-10-24 18:18:00 是否在 2022-11-12 19:45:36 之前？ true
	// 是否在 2022-11-12 19:45:36 之后？ false
	// 还是相等？ false
	fmt.Printf("2022-10-24 18:18:00 是否在 %s 之前？ %v \n 是否在 %s 之后？ %v \n还是相等？ %v\n", s, isBefore, s, isAfter, isEqual)

	sTime := time.Now()
	time.Sleep(time.Second * 3)
	// time.Since(t Time) Duration 【当前时间与时间 t 的时间差（也就是当前时间减去 t 的差）】
	// time.Until(t Time) Duration 【时间 t 与当前时间的差（也就是时间 t 减去当前时间）】
	// 程序开始执行时间为：1666626449 结束执行时间为：1666626452 执行了多长时间：3.003903595s 秒钟
	fmt.Printf("程序开始执行时间为：%v 结束执行时间为：%v 执行了多长时间：%v 秒钟 \n", sTime.Unix(), time.Now().Unix(), time.Since(sTime))
	// 程序在 -3.004024213s 秒钟前执行
	fmt.Printf("程序在 %v 秒钟前执行 \n", time.Until(sTime))
}

func adjustTimeZone() {
	// 转换时间时，指定时区
	var timeParseInTimeZone = func(layout, inputTime, timezone string) string {
		chinaTimezone, _ := time.LoadLocation(timezone) // 指定时区
		// 尽可能的使用 time.ParseInLocation() 方法，因为可以手动指定时区
		// 而少用 time.Parse() 方法，因为默认的时区为 UTC（零时区）
		t, _ := time.ParseInLocation(layout, inputTime, chinaTimezone)
		return time.Unix(t.Unix(), 0).In(chinaTimezone).Format(layout)
	}

	layout := "2006-01-02 15:04:05"
	inputTime := "2022-10-24 21:56:59"

	// 你本机系统得有你设置的时区才行，不然也会走的你系统默认设定的时区
	t1 := timeParseInTimeZone(layout, inputTime, "Asia/Shanghai")
	t2 := timeParseInTimeZone(layout, inputTime, "Asia/Chongqing")
	t3 := timeParseInTimeZone(layout, inputTime, "PRC")              // 中华人民共和国
	t4 := timeParseInTimeZone(layout, inputTime, "Asia/Singapore")   // 新加坡(UTC+08:00)
	t5 := timeParseInTimeZone(layout, inputTime, "Asia/Tokyo")       // 东京(UTC+09:00)
	t6 := timeParseInTimeZone(layout, inputTime, "Etc/GMT")          // 协调世界时(UTC+00:00)
	t7 := timeParseInTimeZone(layout, inputTime, "Pacific/Honolulu") // 夏威夷(UTC-10:00)
	fmt.Printf("Asia/Shanghai ====> %s \n", t1)
	fmt.Printf("Asia/Chongqing ====> %s \n", t2)
	fmt.Printf("PRC ====> %s \n", t3)
	fmt.Printf("Asia/Singapore ====> %s \n", t4)
	fmt.Printf("Asia/Tokyo ====> %s \n", t5)
	fmt.Printf("Etc/GMT ====> %s \n", t6)
	fmt.Printf("Pacific/Honolulu ====> %s \n", t7)

	// 不设置时区时，默认走的 UTC （零时区）
	tt, _ := time.Parse(layout, inputTime)
	fmt.Println(time.Unix(tt.Unix(), 0).Format(layout))
}
