package warnning

import (
	"gfx_project/common"
	"gfx_project/logsystem"
	"strconv"
	"strings"
)

//isTool 参数是否是辅助判断 需要使用mp 记录
func judgeSecLink(log []*common.CompositeData, mp map[string]bool ,date string , isTool bool ) (sp []*common.Warning , errDis int) {
	sp = make([]*common.Warning , 0)
	sp , errDis = getSecLinkFault(log , mp , isTool )
	return
}

//key 每一个二级链路
//返回值 每个错误信息
//错误的二级链路map  key = v[0].PonPort + v[0].PrimaryBeamSplitter + v[0].TwoStageSpectroscope
func getSecLinkFault(log []*common.CompositeData , mp map[string]bool , isTool bool)(sp []*common.Warning , errDis int ) {

	sp = make([]*common.Warning , 0)
	tmp := getSecondSpllData(log)
	for _ , v := range  tmp {
		//排除同一pon下的一级故障
		if v[0] == nil || mp[v[0].PonPort + v[0].PrimaryBeamSplitter] || mp[v[0].PonPort]  {
			continue
		}
		min , max ,ishow , dis:= judgeEverySecondLinkFault(v)
		errDis = dis
		//更新上一次错误信息
		if ishow {
			var obj *common.Warning
			data  :=  GetLastSingleWarnData(SECOND_LEVEL_LINK , v[0].PonPort , v[0].Ip,v[0].OLTName )
			if data == nil {
				obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,SECOND_LEVEL_LINK, dis , min, max ,1, v[0].FinalTime ,
					10 , 0 ,len(v),v[0].ONUAccount,v[0].RecvOptPower,
					v[0].PrimaryBeamSplitter,v[0].TwoStageSpectroscope,v[0].PriPon,v[0].TwoPon,0)
			}else {
				lastime := data.FaultHappenedTime + 10
				num := data.FaultNum + 1
				obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,SECOND_LEVEL_LINK , data.ErroeDiscrde, min, max ,num, v[0].FinalTime ,lastime , 0 ,
					len(v) ,v[0].ONUAccount,v[0].RecvOptPower,v[0].PrimaryBeamSplitter,v[0].TwoStageSpectroscope,v[0].PriPon,v[0].TwoPon,0)
			}

			sp = append(sp ,obj)
			//保存错误的二级链路 的key值到map  ,用于下面的二级分光器错误判断
			if !isTool {
				mp[v[0].PonPort + v[0].PrimaryBeamSplitter + v[0].TwoStageSpectroscope] = true
			}
		}
	}
	return
}

//二级链路
//判断二级链路故障
//判断标准：下面 90% 以上的onu 用户出现问题
func judgeEverySecondLinkFault(data []*common.CompositeData) (minre float64 , maxre float64 , isFault bool , errorDis int){
	var normal int
	var unormalSuspend int
	var unormalbad int
	for i:= 0; i< len(data) ; i++ {
		str := strings.Replace(data[i].RecvOptPower, "\r", "", -1)
		num ,err  :=  strconv.ParseFloat(str, 64)
		if err != nil {
			logsystem.Gfxlog.Err("juedeEveryTwoState strconv ParseFloat error" , err)
			num = float64(0)
			data[i].RecvOptPower = "0"
		}
		errType , ok := GetRecvAbnoralType(num)
		if   ok && errType == CHild_ONU_OUTLINE {
			unormalSuspend++
		}else if ok && (errType == Child_Weak_light  || errType == Child_Extremely_Weak_light ) {
			unormalbad++
		} else if ok == false {
			normal++
		}

		if num < minre {
			minre = num
		}else if num > maxre {
			maxre = num
		}
	}


	probability := float64( float64(unormalSuspend + unormalbad) / float64(unormalSuspend + unormalbad + normal))
	//fmt.Println("judgeEverySecondLinkFault 当前的所有数据 =  === " , unormalSuspend , unormalbad , normal , probability)
	if probability >= 0.9 {
		isFault = true
	}

	if unormalSuspend == (unormalSuspend + unormalbad + normal) {//全部出现下降 就判断光路中断
		errorDis =  Child_Optical_Interruption
		return
	}
	errorDis = Child_Deter_Light_Path
	return
}

//获取二级链路  二级分光器的判断数据
func getSecondSpllData(log []*common.CompositeData) (mp map[string] []*common.CompositeData)  {
	mp = make(map[string] []*common.CompositeData)
	for _, v := range  log {
		hash := v.PonPort + v.PrimaryBeamSplitter + v.TwoStageSpectroscope
		_, ok := mp[hash]
		if !ok {
			mp[hash] = make([]*common.CompositeData , 0)
		}
		mp[hash] = append(mp[hash] ,v)
	}

	//fmt.Printf("getSecondSpllData长度 ===== %v \n" , len(mp))
	return
}

//根据string判断是否是二级链路故障
func judgeIsSecLinkErrorByHash(log []*common.CompositeData) (isTrue bool ,min  float64 ,max  float64 ,errDis int ) {
	data ,errDis  := judgeSecLink(log , nil , "" , true)
	//return len(data) > 0 , data[0]
	isTrue = len(data) > 0
	if isTrue {
		min = data[0].DownMin
		max = data[0].DownMax
	}
	return
}