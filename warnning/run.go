package warnning

import (
	"fmt"
	"gfx_project/common"
	db "gfx_project/datalayer"
)

//父目录警告
/*
	/**
	中断：光功率为-40或者0
	劣化：光功率下降【如-26dbm下降到-30dbm】

	一级链路故障【中断/劣化】：所有二级链路出现故障【中断或者劣化】；
	一级分光器故障：二级链路故障大于等于2个【不区分中断或者劣化】；
	二级链路故障【中断/劣化】：这个二级分光器下面90%以上的用户光功率出现中断或者劣化；
	二级分光器故障：这个二级分光器大于等于2个以上的用户光功率下降【不区分中断或者劣化】；
	尾纤或者连接器故障：同一时刻，整个PON端口下面只有一个用户出现光功率劣化；
	用户下线：同一时刻，整个PON端口下面只有一个用户出现光功率为0或者-40；
*/

//进行报警检测
func RunPlice(log []*common.CompositeData ,date string ) {

	//从错误的等级逐层判断
	name , ok  := db.GetFaultTableName(date)

	if !ok {
		tname := "falsealarm" + date
		err := db.CreateFaultTable(tname)
		if err != nil {
			return
		}
	}

	namearv , ok := db.GetArvTableName(date)
	if !ok {
		db.CreateAvrTime(namearv)
	}

	//1 一级链路判断
	sp1, mp := judegFirstLink(log, date)

	fmt.Printf("judegFirstLink sp1 = %v\n" , len(sp1) )
	//fmt.Printf("CreateFaultTable mp = %v\n" , mp)
	//2 一级分光器判断
	//
	sp2 := judegFirstSplliter(log , mp)
	fmt.Printf("judgeFirstSplliter sp2 = %v\n" ,len(sp2) )

	//3 二级链路判断
	sp3 ,_ := judgeSecLink(log,mp  ,date , false)
	fmt.Printf("judgeFirstSplliter sp3 = %v\n" , len(sp3) )

	// 4. 二级分光器判断
	sp4 := judgeFuncSecondSp(log,mp ,date)
	fmt.Printf("judgeFirstSplliter sp4 = %v\n" ,len(sp4) )

	//5. 个人判断
	sp5 := judgeSingleUser(log , mp,date)
	fmt.Printf("judgeFirstSplliter sp5 = %v\n" , len(sp5))

	//更新错误状态
	spl := [] []*common.Warning{sp1,sp2,sp3,sp4,sp5}
	tempMp := make(map[int] []*common.Warning)
	for i := 0 ;i < len(spl); i++ {
		SetAllFaultData(spl[i] ,tempMp)
	}

	name = getUsefulName(date)

	getResumeData(tempMp ,LastTimeAllFaultData, name ,date)

	for i:= 0; i< len(spl) ; i++ {
		for j := 0; j < len(spl[i]);j++ {
			db.SetFaultDbToDBase(name , spl[i][j])
		}
	}

	//TODO:更新上一次的错误信息
	resetMap(spl)
}

func resetMap(spl [] []*common.Warning) {
	LastTimeAllFaultData = nil
	LastTimeAllFaultData = make(map[int] []*common.Warning)

	for i := 0 ;i < len(spl); i++ {
		SetAllFaultData(spl[i] ,LastTimeAllFaultData)
	}
}




