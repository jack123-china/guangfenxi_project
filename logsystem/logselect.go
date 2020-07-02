package logsystem

import (
	"log/syslog"
	"log"
)

type customLog struct {
	Info func(i ...interface{})
	Err func(i ...interface{})
	Emerg func(i ...interface{})
}

func InitLog(system int) {
	sys , err := syslog.Dial("", "", syslog.LOG_LOCAL0, "GFX ANALYZER")
	sysLog = sys
	if err !=nil {
		log.Fatal(err)
		return
	}
	logs[0].Info = logInfo
	logs[0].Err = logErr
	logs[0].Emerg = logEmerg

	logs[1].Info = log.Println
	logs[1].Err = log.Println
	logs[1].Emerg = log.Fatal

	if system != 0 && system != 1 {
		panic("system != 0 && system != 1")
	}
	Gfxlog = logs[system]
}


func logInfo(i ...interface{}) {
	err := sysLog.Info(i[0].(string))
	if err != nil {
		log.Println(err)
	}
}

func logErr(i ...interface{}) {
	err := sysLog.Err(i[0].(string))
	if err != nil {
		log.Println(err)
	}
}

func logEmerg(i ...interface{}) {
	err := sysLog.Emerg(i[0].(string))
	if err != nil {
		log.Println(err)
	}
}
