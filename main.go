package main

import (
	"fmt"
	"gfx_project/conf"
	"gfx_project/datacollection"
	"gfx_project/monitor"
	"gfx_project/warnning"
	db "gfx_project/datalayer"
	"gfx_project/logsystem"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initDb() (err error) {
	dns := "shijiadong:gfx33561775@tcp(192.168.40.182:3306)/gfx_20200203?charset=utf8"
	err = db.Init(dns)
	if err != nil {
		fmt.Printf("初始化数据库错误 = %v \n", err)
		return
	}
	return
}

func CloseSingel(ch chan int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT,syscall.SIGKILL)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT,syscall.SIGKILL:
			db.Close()
			ch <- 1
			fmt.Printf("捕捉到信号 \n")
		case syscall.SIGHUP:
		default:
			time.Sleep(time.Second)
			fmt.Printf("无信号 \n")
		}
	}
}

func main() {
	closeChan := make(chan int, 0)  //退出管道

	err := conf.ReadFile()  //读取配置
	if err != nil {
		fmt.Println(err)
	}
	_ = initDb()   //初始化数据库

	go CloseSingel(closeChan)


	logsystem.InitLog(1)  //初始化日志

	warnning.InitWarn(datacollection.CurrentDate)

	go monitor.FileMonitor("/home/tmp")

	logpath := "/home/test/a_ft_se_iag_lowlight_user_10m_20200203.csv"

	sp := "/home/test/spliter.txt"
	dis := "/home/test/d_rnt_rm_device_onu_d010"

	datacollection.InitAllData(logpath, sp ,dis)

	select {
		case <-closeChan:
			fmt.Println("退出 main 函数")
			return
	}
	return
}
