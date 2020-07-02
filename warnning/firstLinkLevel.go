package warnning

//一级链路
import (
	"gfx_project/common"
)


//各个判断的调用函数
//一级链路判断
func judegFirstLink(log []*common.CompositeData , date string ) (spwaring []*common.Warning,mp map[string]bool){
	firLinkjudge := getFirstLinkData(log)

	//1.2 判断一级链路
	spwaring , mp = getFirstFault(firLinkjudge)

	return
}

//获得一级链路故障所需要的数据
func getFirstLinkData(log []*common.CompositeData) ( map[string] map[string][]*common.CompositeData){
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


func getFirstFault(tmp map[string]map[string][]*common.CompositeData)(sp []*common.Warning , mp map[string]bool) {
	mp = make(map[string]bool)


	for k , v := range  tmp {
		//每个key 对应的数组  都是二级链路故障
		if len(v) == 0 {
			continue
		}

		var normal int
		var unormal int
		var min float64
		var max float64
		var errDis int
		var ip string
		var otlname string
		var finalTime string
		var account string
		var recv string
		var pri string
		var two string
		var priPon string
		var twoPon string
		for _  ,val := range  v {
			if len(val) == 0 {
				continue
			}
			//fmt.Printf(" 二级map key = %v \n ", key )
			//fmt.Printf("当前二级map长度= %v \n " , len(val) )
			istrue , minC , maxC  , dis := judgeIsSecLinkErrorByHash(val)
			errDis = dis
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
			pri = val[0].PrimaryBeamSplitter
			two = val[0].TwoStageSpectroscope
			priPon = val[0].PriPon
			twoPon = val[0].TwoPon
		}

		//fmt.Printf("当前出现故障的二级链路数 = %v , 当前没有出现故障的二级链路书 = %v \n", unormal ,normal )
		if normal == 0 { //全部二级链路故障 就判断为一级链路故障
			data :=  GetLastSingleWarnData(FIRST_LEVEL_LINK , k, ip,otlname)
			var obj *common.Warning
			if data == nil {//没有当前数据
				obj = NewWarning(otlname,k ,ip ,FIRST_LEVEL_LINK , errDis, min, max ,1, finalTime ,10 ,
					0 ,len(v), account,recv ,pri , two,priPon,twoPon ,0)
			}else {
				// 创建错误类型的对象
				lastime := data.FaultHappenedTime + 10
				num := data.FaultNum + 1
				obj = NewWarning(otlname,k,ip,FIRST_LEVEL_LINK , errDis, min, max ,num,
					finalTime,lastime , 0 ,len(v) ,account,recv ,pri , two,priPon,twoPon ,0)
			}
			sp = append(sp ,obj)

			mp[k] = true  //统一pon下的所有错误 都不再分析
		}
	}

	return
	/*for _ ,v := range  tmp {
		if len(v) == 0 {
			continue
		}
		errType , errorDis, min, max , ok := getLinkFault(v)
		//fmt.Printf("getFirstFault data errType = %v , errorDis = %v ,min = %v , max = %v , ok = %v \n " , errType , errorDis, min, max , ok)

		if !ok {
			continue
		}

		data :=  GetLastSingleWarnData(errType , v[0].PonPort , v[0].Ip,v[0].OLTName )
		var obj *common.Warning
		if data == nil {//没有当前数据
			obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,errType , errorDis, min, max ,1, v[0].FinalTime ,10 , 0 ,len(v), v[0].ONUAccount,v[0].RecvOptPower,id )
		}else {
			// 创建错误类型的对象
			lastime := data.FaultHappenedTime + 10
			num := data.FaultNum + 1
			obj = NewWarning(v[0].OLTName,v[0].PonPort,v[0].Ip,errType , errorDis, min, max ,num, v[0].FinalTime ,lastime , 0 ,len(v) ,v[0].ONUAccount,v[0].RecvOptPower ,id )
		}
		sp = append(sp ,obj)
		//保存错误的一级链路 的key值到map  ,用于下面的一级链路错误判断
		mp[v[0].PonPort + v[0].PrimaryBeamSplitter] = true
	}
	return*/
}
/*
//返回一级链路故障  以及故障描述  最小下降 最大下降
//检查每个一级链路
func getLinkFault(log []*common.CompositeData)  (parentError int , childError int , minre float64 , maxre float64 , iaFault bool){
	logsystem.Gfxlog.Info("==========call getLinkFault============ ")
	var normal int
	var unormal int
	//var suspend int  //中断

	/*for i:= 0; i < len(log) ; i++ {
		str := strings.Replace(log[i].RecvOptPower, "\r", "", -1)
		num , err := strconv.ParseFloat(str, 64)
		if err != nil {
			logsystem.Gfxlog.Err("getLinkFault strconv ParseFloat error = ",err)
			continue
		}
		errType , ok := GetRecvAbnoralType(num)
		fmt.Printf("getLinkFault GetRecvAbnoralType num = %v , errortype = %v , ok = %v \n" , num , errType , ok )

		if ok && (errType == Child_Weak_light || errType == Child_Extremely_Weak_light || errType == CHild_ONU_OUTLINE ) {
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
	judgeIsSecLinkErrorByHash(log)

	/*probability := float64( float64(unormal) / float64(unormal + normal))
	fmt.Printf("一级链路故障的概率 = %v  ,unormal = %v , normal = %v  \n" , probability, unormal , normal)
	if probability > 0.9 {
		parentError = FIRST_LEVEL_LINK
		childError = Child_Deter_Light_Path
		iaFault = true
	}
	return
}
*/


