package warnning

import (
	"gfx_project/common"
	db "gfx_project/datalayer"
)

//mp1 这一次的错误数据
//mp2 上一次的错误数据
//检索上一次的错误 对比这一次
func getResumeData(mp1 ,mp2 map[int] []*common.Warning ,name string , date string) {
	obj := make([]*common.Warning , 0)

	for k ,v := range mp2{
		if mp1[k] == nil {
			obj = append(obj , v...)
		}

		for i := 0; i < len(v) ;i++ {
			isBool := false
			for j := 0; j < len(mp1[k]) ;j++ {
				var hash string
				var hashmp string
				if v[i].ErrorType == FIRST_LEVEL_LINK || v[i].ErrorType == FIRST_SPLLITER {
					hash = v[i].Ponport + v[i].PrimaryBeamSplitter
					hashmp = mp1[k][j].Ponport + mp1[k][j].PrimaryBeamSplitter
				}else if v[i].ErrorType == SECOND_LEVEL_LINK || v[i].ErrorType == SECOND_SPLLITER {
					hash = v[i].Ponport + v[i].PrimaryBeamSplitter + v[i].TwoStageSpectroscope
					hashmp = mp1[k][j].Ponport + mp1[k][j].PrimaryBeamSplitter+ mp1[k][j].TwoStageSpectroscope
				}else if v[i].ErrorType == PIGTAIL_OR_CONN {
					hash = v[i].Ponport + v[i].PrimaryBeamSplitter + v[i].TwoStageSpectroscope + v[i].ONUAccount
					hashmp = mp1[k][j].Ponport + mp1[k][j].PrimaryBeamSplitter + mp1[k][j].TwoStageSpectroscope + mp1[k][j].ONUAccount
				}

				if hash == hashmp {
					isBool = true
				}
			}

			if !isBool {
				obj = append(obj, v[i])
			}
		}
	}

	newTips := db.GetMaxTimeTips("tableLog" + date)
	for _,v := range obj {
		if v.ErrorType == PIGTAIL_OR_CONN {
			db.SetFaultStateSingle(name , v,newTips)
		}else {
			db.SetFaultStateOutOfSingle(name , v,newTips)
		}
	}
}
