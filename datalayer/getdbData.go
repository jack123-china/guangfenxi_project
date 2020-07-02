package datalayer

import (
	"fmt"
	"gfx_project/common"
	"gfx_project/util"
)



//获得分光器数据
func GetSpliterData(port , account string) []*common.Splitter{
	//fmt.Printf("GetSpliterData 的参数 = %v,   第二个 = %v\n"  ,port,account)

	sqlstr := `select * from splliterTable where ponport=? and onuaccount=?`
	obj := make([]*common.Splitter , 0)
	err := DB.Select(&obj, sqlstr,port, account)
	if err != nil {
		return nil
	}

	return obj
}

func GetDistanceData(wd string ,tablename string) []*common.OptDistance {
	sqlstr := `select * from ` +" %s "+` where onupasswd=?`

	sqlstr = fmt.Sprintf(sqlstr , "tableDis"+tablename)
	obj :=make([]*common.OptDistance , 0)
	err := DB.Select(&obj, sqlstr,wd)
	if err != nil {
		return nil
	}

	return obj
}


func GetLogDataByUser(data *common.UserGroup) (obj []*common.CompositeData , err error ) {
	sql := `select * from ` + " %s " + `where  ponport=? and primarybeamsplitter=? and twostagespectroscope=? and onuaccount=? order by tipstime desc limit 1`

	//judgeTableIsExists
	name :="tableLog"+ util.GetToday()
	err = judgeTableIsExists(name)
	if err != nil {
		name = "tableLog"+ util.GetYesterday()
		err = judgeTableIsExists(name)
		if err != nil {
			return
		}
	}

	sql = fmt.Sprintf(sql ,name)
	obj = make([]*common.CompositeData , 0)
	err = DB.Select(&obj, sql,data.PonPort , data.PrimaryBeamSplitter ,data.TwoStageSpectroscope , data.ONUAccount)
	if err != nil {
		return
	}

	return
}


func GetWarnFaultData(tablename string , tmp *common.Warning) int {
	var num []int
	sql := `select faultnum from ` + " %s " + `where faultstate=0 and  ponport=? and ip=? and errortype=? and errdiscribe=? and account=?`
	sql = fmt.Sprintf(sql,tablename)
	_ = DB.Select( &num,sql , tmp.Ponport , tmp.Ip , tmp.ErrorType ,tmp.ErroeDiscrde,tmp.ONUAccount)
	if len(num )> 0 {
		return num[0]
	}
	return 0
}