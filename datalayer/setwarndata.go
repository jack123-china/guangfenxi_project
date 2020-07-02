package datalayer

import (
	"fmt"
	"gfx_project/common"
	"gfx_project/logsystem"
	"gfx_project/util"
)

//设置一级链路 ,一级分光器 故障数据
func SetFaultDbToDBase(name string , data *common.Warning) error {
	sql := `insert into ` + ` %s ` + `(otlname,ponport,ip,errortype,errdiscribe,downmin,downmax,falutnum,finalhappened,
						heppenedtime,faultstate,effectnum,account,recvoptpower,primarybeamsplitter,twostagespectroscope,
						pripon,twopon,tipstime , isinglerror) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	sql = fmt.Sprintf(sql ,name)
	t := util.GetTimeTips()

	_, err := DB.Exec(sql, data.OLTName, data.Ponport, data.Ip , data.ErrorType, data.ErroeDiscrde,data.DownMin,data.DownMax,data.FaultNum,
		data.FaultFinalHappened,data.FaultHappenedTime,data.FaultState,data.EffectNum,data.ONUAccount,data.RecvOptPower,
		data.PrimaryBeamSplitter,data.TwoStageSpectroscope, data.PriPon,data.TwoPon ,t ,data.IsSingleError)
	if err != nil {
		logsystem.Gfxlog.Err("SetFirstLink insert into table error" , err)
		return err
	}
	return nil
}

func SetFaultStateOutOfSingle(tablename string ,tmp *common.Warning ,tips string )  {
	sqlstr := `update ` + " %s " + `faultstate set faultstate=1 
							where faultstate=0 and ponport=? and  ip=? and errortype=? and errdiscribe=? and tipstime<>?`

	sqlstr = fmt.Sprintf(sqlstr , tablename)
	_, err := DB.Exec(sqlstr , tmp.Ponport , tmp.Ip , tmp.ErrorType ,tmp.ErroeDiscrde,tips)
	if err != nil {
		logsystem.Gfxlog.Err("SetFaultState error or not exists tables")
		return
	}
	return
}

func SetFaultStateSingle(tablename string ,tmp *common.Warning ,tips string )  {
	sqlstr := `update ` + " %s " + `faultstate set faultstate=1 
							where faultstate=0  and ponport=? and  ip=? and errortype=? and errdiscribe=? and account=? and tipstime<>?`

	sqlstr = fmt.Sprintf(sqlstr , tablename)
	_, err := DB.Exec(sqlstr , tmp.Ponport , tmp.Ip , tmp.ErrorType ,tmp.ErroeDiscrde,tmp.ONUAccount ,tips)
	if err != nil {
		logsystem.Gfxlog.Err("SetFaultState error or not exists tables")
		return
	}
	return
}

