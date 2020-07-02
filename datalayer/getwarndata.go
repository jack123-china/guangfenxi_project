package datalayer

import (
	"fmt"
	"gfx_project/common"
	"gfx_project/logsystem"
	"gfx_project/util"
)


//获得一级链路 的错误数据
//一级链路的故障 通过 pon口 跟 错误等级查询 取最新的一条
func GetFirstLinkData(ponport string,tablename string) (re *common.Warning , err error){
	sql := `select * ` + " %s " + `where ponport =? and FaultState = ? order by finalhappened limit 1`

	tmp := make([]*common.Warning, 0)
	err = DB.Select(&tmp,sql, ponport, 0)
	if err != nil {
		logsystem.Gfxlog.Err("GetFirstLinkData select error")
		return
	}

	if len(tmp) > 0 {
		re = tmp[0]
	}
	return
}

//date 为日志上的日期
//false 时候创建
func GetFaultTableName(date string) (string , bool ) {
	name := "falsealarm" + date
	err := judgeTableIsExists(name)
	if err != nil {
		name = "falsealarm" + util.GetYesterday()
		return name ,false
	}
	return name , true
}

//平均时间表
func GetArvTableName(date string) (string , bool ) {
	name := "avr_time" + date
	err := judgeTableIsExists(name)
	if err != nil {
		return name ,false
	}
	return name , true
}

//获得最新的 报警数据
func GetNewestFaultData(tablename ,timeTips string ) ([] *common.Warning, error){
	tmp := make([]*common.Warning, 0)
	if timeTips == "" {
		return tmp , nil
	}

	sql := `select * from ` + " %s " + `where tipstime=?`
	sql = fmt.Sprintf(sql ,tablename)

	err := DB.Select(&tmp,sql ,timeTips)
	if err != nil {
		logsystem.Gfxlog.Err("GetNewestFaultData is not exist" )
		return nil ,err
	}
	return tmp , nil
}

