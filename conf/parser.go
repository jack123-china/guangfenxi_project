package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gfx_project/logsystem"
)

func ReadFile()  (err error){
	cfg, err := ini.Load(__CONF__)
	if err != nil {
		return
	}
	cf := TestConf{}
	cf.CurLogType,_ = cfg.Section("server").Key("server_curlogtype").Int()
	fmt.Printf("cf.CurrnetLogType = %v\n", cf.CurLogType)
	logsystem.InitLog(cf.CurLogType)
	return
}
