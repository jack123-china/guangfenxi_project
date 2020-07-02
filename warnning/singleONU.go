package warnning

import (
	"gfx_project/common"
	"gfx_project/logsystem"
	"strconv"
	"strings"
)

func judgeSingleUser(log []*common.CompositeData,mp map[string]bool ,date string) (sp []*common.Warning) {

	sp = singleONUser(getSingleUserData(log) ,mp )
	getUsefulName(date)

	return
}

func singleONUser(log map[string]*common.CompositeData ,mp map[string]bool) (sp []*common.Warning){

	for _ ,v := range log{
		str := strings.Replace(v.RecvOptPower, "\r", "", -1)
		num ,err := strconv.ParseFloat(str, 64)

		isError := 0
		if mp[v.PonPort + v.PrimaryBeamSplitter + v.TwoStageSpectroscope] || mp[v.PonPort + v.PrimaryBeamSplitter] || mp[v.PonPort] {
			isError = 1
		}

		if err != nil {
			logsystem.Gfxlog.Err("judgeSingleUser strconv.ParseFloat ")
			continue
		}
		errType , ok := GetRecvAbnoralType(num)
		//fmt.Printf("===========判断 单人 数据 =errType = %v,ok = %v \n ", errType , ok)

		if ok {
			//1.创建错误对象
			data  :=  GetLastSingleWarnData(PIGTAIL_OR_CONN , v.PonPort ,v.Ip,v.OLTName)

			// 创建错误类型的对象
			var obj *common.Warning
			if data != nil {
				lastime := data.FaultHappenedTime + 10
				num := data.FaultNum + 1
				obj = NewWarning(v.OLTName,v.PonPort,v.Ip,PIGTAIL_OR_CONN , errType, 0 , 0 ,num,
					v.FinalTime ,lastime , 0 ,1,v.ONUAccount,v.RecvOptPower ,
					v.PrimaryBeamSplitter , v.TwoStageSpectroscope, v.PriPon,v.TwoPon,isError)
			}else {
				obj = NewWarning(v.OLTName,v.PonPort,v.Ip,PIGTAIL_OR_CONN , errType, 0 , 0 ,1,
					v.FinalTime ,10 , 0 ,1,v.ONUAccount,v.RecvOptPower ,
					v.PrimaryBeamSplitter , v.TwoStageSpectroscope, v.PriPon,v.TwoPon,isError)
			}
			sp = append(sp ,obj)
		}
	}
	return
}

//获得单个用户
func getSingleUserData(log []*common.CompositeData)  (mp map[string] *common.CompositeData)  {
	mp = make(map[string] *common.CompositeData)
	for _, v := range  log {
		hash := v.ONUAccount + v.ONUPassWd
		mp[hash] = v
	}

	return
}



