package datalayer

import (
	"fmt"
	"gfx_project/logsystem"

)

//falsealarm
func judgeTableIsExists(name string) error {
	fmt.Printf("call  judgeTableIsExists \n")

	slq := `select COUNT(id) from ` + ` %s ` + `where id=1`
	slq = fmt.Sprintf(slq , name)
	var count int
	err := DB.Get(&count,slq)

	if err != nil {
		logsystem.Gfxlog.Err("current table is not exist" , err)
		return err
	}
	return nil
}



func IsExitsTable(name ,column string ) bool {
	logsystem.Gfxlog.Info("call IsExitsTable")

	sql := `SELECT COUNT(` + ` %s ` +  `) FROM` +  " %s "  + ` where id=1`
	sql = fmt.Sprintf(sql ,column,name)
	var num int
	err := DB.Get(&num,sql)
	if num == 0 || err != nil {
		logsystem.Gfxlog.Err("Select num error")
		return false
	}
	return true
}
