package datalayer

import (
	"fmt"
	"errors"
	"gfx_project/common"
	"gfx_project/logsystem"
)

//解析出的最终结果插入数据库
func SetParserDataToDB(data *common.CompositeData ,name string)  (err error) {
	name = "tableLog" + name
	sqlStr := `insert into ` + ` %s ` + `(cityname,otlname,ponport,primarybeamsplitter,twostagespectroscope,
											onuaccount,onupasswd,recvoptpower,sendoptpower,dispontoonu,ip,time,pripon,twopon,tipstime) 
										values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	str := fmt.Sprintf(sqlStr , name)

	_, err = DB.Exec(str, data.CityName,data.OLTName,data.PonPort,data.PrimaryBeamSplitter,data.TwoStageSpectroscope,
		data.ONUAccount,data.ONUPassWd,data.RecvOptPower,data.SendOptPower,data.DistanceFromPONToONU,data.Ip,data.FinalTime,data.PriPon,data.TwoPon,data.TipsTime)
	if err != nil {
		fmt.Printf("SetParserDataToDB data address = %v \n" ,err )
		return
	}
	return
}

//插入分光器表
func InsertIntoSplliterTable(data *common.Splitter)  (err error)  {
	sqlStr := `select * from splliterTable where ponport= ? and onuaccount = ?`
	var d []*common.Splitter
	err = DB.Select(&d, sqlStr, data.PonPort,data.ONUAccount)
	if err != nil {
		fmt.Printf("InsertIntoSplliterTable  %v\n" , err)
		return
	}

	if len(d) > 0{
		logsystem.Gfxlog.Info("currrent spliter data in databases")
		err = errors.New("currrent spliter data in databases")
		return
	}
	sqlStr = `insert into splliterTable (ponport,primarybeamsplitter,
			twostagespectroscope,onuaccount) values (?,?,?,?)`

	_, err = DB.Exec(sqlStr, data.PonPort,data.PrimaryBeamSplitter,data.TwoStageSpectroscope,data.ONUAccount)
	if err != nil {
		return
	}
	return
}

//插入光距表
func InsertIntoDisTable(data *common.OptDistance , name string)  (err error) {
	fmt.Printf("插入光距表 = %v \n" ,name)

	name = "tableDis" + name
	sqlStr := `select * from ` + " %s "+ ` where onupasswd= ? `
	sqlStr = fmt.Sprintf(sqlStr , name)

	var d []*common.OptDistance
	err = DB.Select(&d, sqlStr, data.ONUPassWd)
	if err != nil {
		return
	}

	if len(d) > 0 {
		logsystem.Gfxlog.Info("currrent distancetable data in databases")
		err = errors.New("currrent distancetable data in databases")
		return
	}

	sqlStr = `insert into ` +   " %s " + `(onupasswd,onuupline,
			onudownline,dispontoonu) values (?,?,?,?)`
	sqlStr = fmt.Sprintf(sqlStr , name)

	_, err = DB.Exec(sqlStr, data.ONUPassWd,data.ONUUpLine,data.ONUDownLine,data.DistanceFromPONToONU)
	if err != nil {
		return
	}
	return
}


//预处理插入分光器
func PreInsertSplliter(sp map[string]* common.Splitter)  {
	sqlStr := `insert into splliterTable (ponport,primarybeamsplitter,
			twostagespectroscope,onuaccount,pripon,twopon) values (?,?,?,?,?,?)`

	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("DB.Prepare error = %v \n",err)
		return
	}
	for _ , v := range sp {
		_, err = stmt.Exec(v.PonPort,v.PrimaryBeamSplitter,v.TwoStageSpectroscope,v.ONUAccount,v.PriPon,v.TwoPon)
		if err != nil {
			fmt.Println("DB.Exec error")
			continue
		}
	}

	fmt.Println("PreInsertSplliter finish \n")
}

func PreInsertDis(sp map[string]* common.OptDistance , name string) {
	name = "tableDis" + name
	sqlStr := `insert into ` +   " %s " + `(onupasswd,onuupline,
			onudownline,dispontoonu) values (?,?,?,?)`
	sqlStr = fmt.Sprintf(sqlStr , name)

	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("PreInsertDis DB.Prepare error = %v \n",err)
		return
	}

	for _ , v := range sp {
		_, err = stmt.Exec(v.ONUPassWd,v.ONUUpLine,v.ONUDownLine,v.DistanceFromPONToONU)
		if err != nil {
			fmt.Println("PreInsertDis DB.Exec error")
			continue
		}
	}

	fmt.Println("PreInsertDis finish \n")
}

//改变结果表光距值
func ChangeFinalTableDis(changetable string ,dis map[string]* common.OptDistance) {
	name := "tableLog" + changetable
	sqlStr := `update `+ " %s "+ ` set dispontoonu = ? where onupasswd = ?`
	sqlStr = fmt.Sprintf(sqlStr,name )
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return
	}
	for _ , v := range  dis {
		_, err = stmt.Exec(v.DistanceFromPONToONU , v.ONUPassWd)
		if err != nil {
			continue
		}
	}
}

//更新用户分组信息
func CheckUsergroup(data []*common.UserGroup)  {
	logsystem.Gfxlog.Err("call CheckUsergroup==============")
	for _ , v := range  data {
		slq := `select COUNT(*) from only_usergroup where ponport=? and  onuaccount=?`
		var count int
		err := DB.Get(&count,slq, v.PonPort,v.ONUAccount)
		if err != nil {
			logsystem.Gfxlog.Err("查询用户唯一表出现错误==============" ,err)
			continue
		}

		if(count > 0) {
			logsystem.Gfxlog.Err("用户分组表发现到相同数据==============" ,count)
			continue
		}else{
			slq = `insert into only_usergroup(otlname,address,ponport,primarybeamsplitter,twostagespectroscope,onuaccount,pripon , twopon) values(?,?,?,?,?,?,?,?)`
			_, err = DB.Exec(slq , v.OLTName , v.Address, v.PonPort, v.PrimaryBeamSplitter, v.TwoStageSpectroscope, v.ONUAccount,v.PriPon,v.TwoPon)
			if err != nil {
				logsystem.Gfxlog.Err("插入用户分组表出错")
				continue
			}
		}
	}
}

//
func GetAllUserData() (data []*common.UserGroup) {
	sql := `select * from only_usergroup`
	err := DB.Select(&data , sql)
	fmt.Printf("GetAllUserData GetAllUserData = %v \n" , len(data) )
	if err != nil {
		fmt.Printf("GetAllUserData error = %v \n" , err )
		return
	}
	return
}


