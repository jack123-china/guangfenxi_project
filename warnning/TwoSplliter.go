package warnning

import (
	"gfx_project/common"
	"gfx_project/logsystem"
	"strconv"
	"strings"
)

//二级分光器
func judgeFuncSecondSp(log []*common.CompositeData ,mp map[string]bool ,date string )(sp []*common.Warning) {
	sp = judgeSecSplliter(log,mp )
	return
}

//二级分光器故障
func judgeSecSplliter(log []*common.CompositeData ,mp map[string]bool ) (sp []*common.Warning){
	//获取数据
	tmp := getSecondSpllData(log)
	for k , v := range tmp {
		if v[0] == nil || mp[k]  || mp[v[0].PonPort + v[0].PrimaryBeamSplitter + v[0].TwoStageSpectroscope] ||
			mp[v[0].PonPort ] || mp[v[0].PonPort + v[0].PrimaryBeamSplitter] {
			continue
		}

		min , max , ishow := juedeEveryTwoState(v)
		if ishow {
			data  :=  GetLastSingleWarnData(SECOND_SPLLITER , v[0].PonPort , v[0].Ip,v[0].OLTName )
			var obj *common.Warning
			if data == nil {
				obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,SECOND_SPLLITER , 0, min, max ,1, v[0].FinalTime ,10 ,
					0 ,len(v),v[0].ONUAccount ,v[0].RecvOptPower,v[0].PrimaryBeamSplitter,v[0].TwoStageSpectroscope,v[0].PriPon,v[0].TwoPon,0)
			}else {
				lastime := data.FaultHappenedTime + 10
				num := data.FaultNum + 1
				obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,data.ErrorType , data.ErroeDiscrde, min, max ,num,
					v[0].FinalTime ,lastime , 0 ,len(v),v[0].ONUAccount,v[0].RecvOptPower ,v[0].PrimaryBeamSplitter,v[0].TwoStageSpectroscope,v[0].PriPon,v[0].TwoPon,0)

			}
			sp = append(sp ,obj)
			//logsystem.Gfxlog.Err("二级分光器故障数据 ========================= %v\n" , len(sp))
			mp[v[0].PonPort + v[0].PrimaryBeamSplitter + v[0].TwoStageSpectroscope] = true
		}
	}
	return
}


//判断每个二级分光器
//大于2个就是故障
//判断标准：两个以上用户出现故障
func juedeEveryTwoState(data []*common.CompositeData) (minre float64 , maxre float64 , isFault bool){
	var normal int
	var unormal int
	for i:= 0; i< len(data) ; i++ {
		str := strings.Replace(data[i].RecvOptPower, "\r", "", -1)
		num ,err  :=  strconv.ParseFloat(str, 64)
		if err != nil {
			logsystem.Gfxlog.Err("juedeEveryTwoState strconv ParseFloat error" , err)
			num = float64(0)
			data[i].RecvOptPower = "0"
		}


		errType , ok := GetRecvAbnoralType(num)
		if ok && (errType == Child_Weak_light || errType == Child_Extremely_Weak_light || CHild_ONU_OUTLINE == errType) {
			unormal++
		} else if ok == false {
			normal++
		}

		if num < minre {
			minre = num
		}else if num > maxre {
			maxre = num
		}
	}

	//fmt.Printf("多少个二级分光器损坏 = %d \n" , unormal)
	if unormal >= 2 {
		isFault = true
	}
	return
}


