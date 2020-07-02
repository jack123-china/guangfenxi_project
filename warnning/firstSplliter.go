package warnning

import (
	"fmt"
	"gfx_project/common"
)



//各个判断的调用函数
//一级链路判断
func judegFirstSplliter(log []*common.CompositeData , mp map[string]bool ) (spwaring []*common.Warning){
	firLinkjudge := getFirstSplliterData(log)

	//1.2 判断一级链路
	spwaring = getFirstSplliterFault(firLinkjudge , mp)

	return
}

//获得一级链路故障所需要的数据
func getFirstSplliterData(log []*common.CompositeData) ( map[string] map[string][]*common.CompositeData){
	//string  为 v.PonPort
	//string  为 v.PonPort + v.PrimaryBeamSplitter + v.TwoStageSpectroscope
	mp_pon := make(map[string] map[string][]*common.CompositeData)

	for _ , v := range log {
		hash := v.PonPort + v.PrimaryBeamSplitter + v.TwoStageSpectroscope
		_ , ok := mp_pon[v.PonPort]
		if !ok {
			mp_pon[v.PonPort] = make(map[string][]*common.CompositeData)
		}
		mp_pon[v.PonPort][hash] = append(mp_pon[v.PonPort][hash] , v)
	}
	return mp_pon
}


func getFirstSplliterFault(tmp map[string] map[string][]*common.CompositeData ,mp map[string]bool )(sp []*common.Warning) {
//	mp = make(map[string]bool)

	for k , v := range  tmp {
		//每个key 对应的数组  都是二级链路故障
		if len(v) == 0 || mp[k]{
			continue
		}

		var normal int
		var unormal int
		var min float64
		var max float64
		var ip string
		var otlname string
		var finalTime string
		var account string
		var recv string
		var spllitername string
		var pri string
		var two string
		var priPon string
		var twoPon string
		for _  ,val := range  v {
			if len(val) == 0 {
				continue
			}

			istrue , minC , maxC  , _ := judgeIsSecLinkErrorByHash(val)
		//	errDis = dis
			if istrue {
				unormal++
			}else {
				normal++
			}

			if minC < min {
				min =minC
			}
			if maxC > max {
				max = maxC
			}
			ip = val[0].Ip
			otlname = val[0].OLTName
			finalTime =  val[0].FinalTime
			account =   val[0].ONUAccount
			recv = val[0].RecvOptPower
			spllitername= val[0].PrimaryBeamSplitter
			pri = val[0].PrimaryBeamSplitter
			two = val[0].TwoStageSpectroscope
			priPon = val[0].PriPon
			twoPon = val[0].TwoPon
		}


		//
		fmt.Printf("一级分光器故障的个数 === %v \n" , unormal)
		if unormal >= 2 { //一级分光器故障  分光器故障无错误描述
			data :=  GetLastSingleWarnData(FIRST_SPLLITER , k, ip,otlname)
			var obj *common.Warning
			if data == nil {//没有当前数据
				obj = NewWarning(otlname,k ,ip ,FIRST_SPLLITER , 0, min, max ,1, finalTime ,
					10 , 0 ,len(v), account,recv ,pri, two,priPon,twoPon  ,0)
			}else {
				// 创建错误类型的对象
				lastime := data.FaultHappenedTime + 10
				num := data.FaultNum + 1
				obj = NewWarning(otlname,k,ip,FIRST_SPLLITER , 0, min, max ,num, finalTime,lastime , 0 ,len(v) ,account,recv ,pri,
					two ,priPon,twoPon ,0)
			}
			sp = append(sp ,obj)

			mp[k + spllitername] = true //一级链路故障  pon+ PrimaryBeamSplitter
		}
	}
	return
}
