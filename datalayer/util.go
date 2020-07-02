package datalayer

import  (
	"fmt"
	"gfx_project/common"
	"gfx_project/logsystem"
	"gfx_project/util"
)


func  GetMaxTimeTips(tablename string) string {
	sql := `select * from ` + ` %s ` + ` order by tipstime desc limit 1`

	sql = fmt.Sprintf(sql , tablename)
	fmt.Println(sql)

	obj :=make([]*common.CompositeData , 0)
	err := DB.Select(&obj,sql)
	if err != nil {
		logsystem.Gfxlog.Err("GetMaxTimeTips Select error" , err)
		return ""
	}

	if len(obj) >= 1 {
		return obj[0].TipsTime
	}
	return ""
}


func GetLastTimeTips(tablename string ) string {
	tips := GetMaxTimeTips(tablename )
	return  util.BeforeTenMin(tips)
}