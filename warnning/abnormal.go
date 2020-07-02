package warnning

import (
	"gfx_project/common"
	"gfx_project/logsystem"
)

//单个用户
//判断 接收光功率  强光 弱光 极弱光 下线
//单个用户判断到下线 就是下线
//多个用户（> 1 ）判断到下线 就是中断报错
//劣化就是都变小了 超出正常范围
func GetRecvAbnoralType(data float64) (errorType int ,ok bool ) {
	ok = false
	if data == float64(Stand_Recv_OutLine_1) || data == float64(Stand_Recv_OutLine_2 ){
		//当前ONU 用户下线
		errorType = CHild_ONU_OUTLINE
		ok = true
		return
	}

	if data > float64(Standard_Recv_Raw_up) {
		errorType = Child_Strong_Light
		ok = true
	}else if data <= float64(Standard_Recv_Extremely_Weak_light) {
		//极度弱光
		ok = true
		errorType = Child_Extremely_Weak_light
	}else if data < float64(Standard_Recv_Raw_dpwn) && data >  float64(Standard_Recv_Extremely_Weak_light){
		//弱光
		ok = true
		errorType = Child_Weak_light
	}
	return
}

func GetLastSingleWarnData(errType int, ponport string ,ip string ,oltname string ) (*common.Warning){
	logsystem.Gfxlog.Info("============== call GetLastSingleWarnData ===============")

	mp := LastTimeAllFaultData[errType]

	for i:= 0; i< len(mp); i++ {
		if errType == mp[i].ErrorType && ponport == mp[i].Ponport && ip ==  mp[i].Ip && oltname == mp[i].OLTName {
			return mp[i]
		}
	}

	return nil
}
