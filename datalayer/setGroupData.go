package datalayer

import (
	"gfx_project/common"
	"gfx_project/logsystem"
)

//设置唯一用户表信息
func SetOnlyGroupData(data []*common.UserGroup) (err error){
	for _, v :=range data {
		slq := `select COUNT(ponport ,onuaccount) from only_usergroup where ponport=? and  onuaccount=?`
		var count int
		err = DB.Get(&count,slq)
		if err != nil {
			continue
		}
		if count > 0 {
			continue
		}


		sql := `insert into only_usergroup(otlname,address,ponport,primarybeamsplitter,twostagespectroscope,onuaccount) value(?,?,?,?,?,?) `
		_, err = DB.Exec(sql, v.OLTName ,v.Address , v.PonPort, v.PrimaryBeamSplitter,v.TwoStageSpectroscope,v.ONUAccount)
		if err != nil {
			logsystem.Gfxlog.Err("SetOnlyGroupData insert into table error" , err)
			continue
		}
	}

	return
}

