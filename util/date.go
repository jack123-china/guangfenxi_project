package util

import (
	"errors"
	"gfx_project/logsystem"
	"strings"
	"time"
)

var MemTimeTips string

func GetYesterday() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,-1)
	logDay := yesTime.Format("20060102")
	return logDay
}

func GetToday() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,0)
	logDay := yesTime.Format("20060102")
	return logDay
}

func StringConvToTime(k string ) (re time.Time) {
	const Layout = "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	re ,_ = time.ParseInLocation(Layout,k,loc)
	return
}

func GetAvrToadyName() string {
	//avr_time20200609
	return "avr_time"+GetToday()
}

//时间比较
func CompareTime(time1, time2 string ) bool {
	/*time1 :=  "2015-03-21 09:04:25"//"2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25" */
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil  {
		//处理逻辑
		//fmt.Println( t1.Before(t2))
		logsystem.Gfxlog.Err("CompareTime error1" , err)
		return false
	}

	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err != nil  {
		//处理逻辑
		//fmt.Println( t1.Before(t2))
		logsystem.Gfxlog.Err("CompareTime error2" , err)
		return false
	}
	return t1.Before(t2)
}

func FormatTime(str string) string {
	tmp  := strings.Split(str, " ")
	date , _ := FormatDateString(tmp[0] , "/")
	times , _ := FormatTimeString(tmp[1] , ":")

	//fmt.Printf("FormatTime FormatDateString = %v , FormatTimeString = %v \n" , date ,times)

	return date + " " + times
}

//格式化日期
func FormatDateString(date string,tips string ) (string, error) {
	tmp  := strings.Split(date, tips)
	if len(tmp) < 3 {
		err := errors.New("当前日期字段错误")
		logsystem.Gfxlog.Err("FormatDateString error " , err)
		return "" ,err
	}

	str := ""
	for i := 0 ; i< len(tmp); i++ {
		if i == len(tmp) - 1 {
			str += (tmp[i])
		}else if i == 1 {
			str += ("0" + tmp[i] + "-")
		}else {
			str += (tmp[i] + "-")
		}
	}

	return str , nil
}

//格式化时间
func FormatTimeString(time string ,tips string) (str string,err error) {
	tmp  := strings.Split(time, tips)
	if len(tmp) == 2 {
		str = tmp[0] + ":" + tmp[1] + ":" + "00"
	}else if len(tmp) == 3 {
		str = tmp[0] + ":" + tmp[1] + ":" + tmp[2]
	}else {
		err = errors.New("当前日期字段错误")
	}
	return str , err
}

//5分钟后
func AfterFiveMin(str string) string  {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t , _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	m, _ :=  time.ParseDuration("1m")
	m1 := t.Add(10 * m)
	return m1.Format("2006-01-02 15:04:05")
}


func SetTimeTips(str string ) {
	MemTimeTips = str
}

func GetTimeTips() string {
	return MemTimeTips
}

func BeforeTenMin(str string) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t , _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	m, _ :=  time.ParseDuration("1m")
	m1 := t.Add(-10 * m)
	return m1.Format("2006-01-02 15:04:05")
}
