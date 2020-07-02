package warnning

import (
	"gfx_project/common"
	db "gfx_project/datalayer"
	"gfx_project/logsystem"
)

//key 错误类型
var LastTimeAllFaultData map[int] []*common.Warning  //保存上一次所有的错误信息

func NewWarning(name string , port string , ip string , errortype int , errordis int , min float64 ,max float64,
			faultnum int , finalhanppend string , lastime int , faultstate int , num int,account string ,rec string, pri string ,
			two string , pripon string , twopon string , isSingle int) *common.Warning {
	return &common.Warning{
		OLTName:            name,
		Ponport:            port,
		Ip:                 ip,
		ErrorType:          errortype,
		ErroeDiscrde:       errordis,
		DownMin:            min,
		DownMax:            max,
		FaultNum:           faultnum,
		FaultFinalHappened: finalhanppend,
		FaultHappenedTime:  lastime,
		FaultState:         faultstate,
		EffectNum:          num,
		ONUAccount:         account,
		RecvOptPower: rec ,
		PrimaryBeamSplitter:pri,
		TwoStageSpectroscope:two,
		PriPon:pripon,
		TwoPon:twopon,
		IsSingleError: isSingle,
	}
}

//初始化
func InitWarn(date string) {
	LastTimeAllFaultData = make(map[int] []*common.Warning)
	timeTips := db.GetMaxTimeTips("tableLog" + date)
	data , err := db.GetNewestFaultData(setWarnName(date) , timeTips)
	if err != nil {
		SetAllFaultData(data,LastTimeAllFaultData)
	}
}

func SetAllFaultData(data []*common.Warning ,mp map[int] []*common.Warning) /*(mp map[string] *common.Warning) */{
	logsystem.Gfxlog.Info("call SetAllFaultData")
	for i:=0 ;i < len(data) ;i++ {
		if data[i].ErrorType != FIRST_LEVEL_LINK  && data[i].ErrorType != FIRST_SPLLITER && data[i].ErrorType != SECOND_LEVEL_LINK &&
			data[i].ErrorType != SECOND_SPLLITER && data[i].ErrorType != SECOND_SPLLITER && data[i].ErrorType != PIGTAIL_OR_CONN  {
			logsystem.Gfxlog.Info("SetAllFaultData errortype error")
			continue
		}
		SetVFaultFun(data[i].ErrorType , data[i] , mp)
	}
}

func SetVFaultFun(errtype int ,data *common.Warning ,mp map[int] []*common.Warning) {
	_, ok := mp[errtype]
	if !ok {
		mp[errtype] = make([]*common.Warning , 0)
	}
	mp[errtype] = append(mp[errtype] , data)
}

func setWarnName(date string) string {
	return "falsealarm" + date
}

func getUsefulName(date string ) string {
	name , ok  := db.GetFaultTableName(date)  //进入的时候表就判断创建
	if !ok {
		name = "falsealarm" + date
	}
	return name
}


