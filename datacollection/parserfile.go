//解析 日志文件

package datacollection

import (
	"errors"
	"gfx_project/common"
	"gfx_project/logsystem"
	"gfx_project/util"
	"io"

	"strconv"
	"strings"
)

//光功率
func PareseRawLogFile(path string) (mp_line_spliter map[string]*common.RawLog){
	mp_line_spliter = make(map[string]*common.RawLog )
	//mp_pon =  make(map[string] []*common.RawLog )
	lines ,err  := util.PareseFile(path)
	if err != nil && err != io.EOF {
		logsystem.Gfxlog.Err("PareseRawLogFile read file error ,file error = ， name = ", err, path)
		return
	}

	legth := len(lines)

	for i := 0; i< legth ; i++ {
		obj ,err := NewRawLogObj(lines[i])
		if err != nil {
			continue
		}

		hash :=  obj.PonPort + obj.ONUAccount
		_, ok := mp_line_spliter[hash]
		if ok {
			continue
		}
		mp_line_spliter[hash] = obj
	}
	return
}

//分光器
func PareseSplliter(path string) ( mp_line map[string]*common.Splitter){
	mp_line = make(map[string]*common.Splitter)
	lines ,err  := util.PareseFile(path)
	if err != nil  && err != io.EOF {
		logsystem.Gfxlog.Err("PareseSplliter read file error ,file error = ， name = ", err, path)
		return
	}
	legth := len(lines)

	for i:= 0; i < legth ; i++ {
		obj  , err := NewSplitterObj(lines[i])
		if err != nil {
			continue
		}
		hash := obj.PonPort + obj.ONUAccount
		_, ok := mp_line[hash]
		if ok {
			logsystem.Gfxlog.Err("PareseSplliter  NewSplitterObj , hash = obj.PonPort + obj.ONUAccount is eixsts =  ",hash )
			continue
		}
		mp_line[hash] = obj
	}
	return
}

//光距
func ParserDistance(path string) (mp_line map[string]*common.OptDistance){
	mp_line = make(map[string]*common.OptDistance)
	lines ,err  := util.PareseFile(path)
	if err != nil   && err != io.EOF{
		logsystem.Gfxlog.Err("ParserDistance read file error ,file error = ， name = ", err, path)
		return
	}
	legth := len(lines)

	for i:= 0; i < legth ; i++ {
		obj := NewOptDistanceObj(lines[i])
		if obj == nil {
			continue
		}
		hash := obj.ONUPassWd
		_, ok := mp_line[hash]
		if !ok {
			mp_line[hash] = obj
		}
	}
	return
}

func NewRawLogObj(data string) (*common.RawLog , error) {
	kov := strings.Split(data, "|")
	if len(kov) != 21 {
		logsystem.Gfxlog.Info("日志异常。。。。字段个数= \n" ,data)
		return  nil,errors.New("日志异常")
	}
	var num string
	if kov[20] == "" {
		num = "0"
	}else {
		num = kov[20]
	}
	tmptime := util.FormatTime(kov[0])
	//fmt.Printf("NewRawLogObj tmptime = %v\n" , tmptime)
	obj := &common.RawLog{
		PonPort:kov[8],
		Time: kov[0],
		CityName:kov[1],
		OLTName : kov[7],
		OLTMgrIp: kov[6],
		OUNumber :kov[12],
		ONUAccount: kov[15],
		ONUPassWd: kov[14],
		RecvOptPower:num,
		SendOptPower:kov[19],
		GateWayRunTime:kov[16],
		TipsTime:tmptime,
	}

	obj.CPUOccupy , _ = strconv.Atoi(kov[17])
	obj.MemOccupy , _ = strconv.Atoi(kov[18])

	return obj , nil
}

func NewSplitterObj(data string) (*common.Splitter ,error){
	kov := strings.Split(data, "|")
	str := strings.Split( kov[8], ",")
	if len(str) < 4 {
		logsystem.Gfxlog.Info("分光器信息异常。。。。")
		return nil ,errors.New("分光器信息异常")
	}

	obj := &common.Splitter{
		PonPort : str[0],
		PrimaryBeamSplitter  :GroupTopString(str[1]),
		PriPon:GroupTailString(str[1]),
		TwoStageSpectroscope :GroupTopString(str[2]),
		TwoPon:GroupTailString(str[2]),
		ONUAccount : kov[3],
	}

	return obj , nil
}

func NewOptDistanceObj(data string) *common.OptDistance {
	kov := strings.Split(data, "|")
	if len(kov) < 9 {
		return nil
	}
	obj:= &common.OptDistance{
		ONUPassWd:             kov[6],
		ONUUpLine:             kov[7],
		ONUDownLine:           kov[8],
		DistanceFromPONToONU:  kov[9],
	}
	return obj
}

//数据组合
func NewCompositeData(log *common.RawLog ,split *common.Splitter ,dis *common.OptDistance)  *common.CompositeData{
	return &common.CompositeData{
		CityName:             log.CityName,
		OLTName:              log.OLTName,
		Ip: 				  log.OLTMgrIp,
		PonPort:              log.PonPort,
		PrimaryBeamSplitter:  split.PrimaryBeamSplitter,
		PriPon:				  split.PriPon,
		TwoStageSpectroscope: split.TwoStageSpectroscope,
		TwoPon:               split.TwoPon,
		ONUAccount:           log.ONUAccount,
		ONUPassWd:            log.ONUPassWd,
		RecvOptPower:         log.RecvOptPower,
		SendOptPower:         log.SendOptPower,
		DistanceFromPONToONU: dis.DistanceFromPONToONU,
		FinalTime: log.Time,
		TipsTime:log.TipsTime,
	}
}

func NewConPositeLog(log *common.CompositeData , time string , tips string )  *common.CompositeData{
	return &common.CompositeData{
		CityName:             log.CityName,
		OLTName:              log.OLTName,
		Ip: 				  log.Ip,
		PonPort:              log.PonPort,
		PrimaryBeamSplitter:  log.PrimaryBeamSplitter,
		PriPon:				  log.PriPon,
		TwoStageSpectroscope: log.TwoStageSpectroscope,
		TwoPon:               log.TwoPon,
		ONUAccount:           log.ONUAccount,
		ONUPassWd:            log.ONUPassWd,
		RecvOptPower:         "0",
		SendOptPower:         "0",
		DistanceFromPONToONU: log.DistanceFromPONToONU,
		FinalTime: time ,
		TipsTime:tips,
	}
}


func NewUserGroup(data *common.RawLog ,pri string , twostate string ,priPon string , twoPon string ) (*common.UserGroup){
	obj := &common.UserGroup{
		Address: data.CityName,
		OLTName:data.OLTName,
		PrimaryBeamSplitter  :pri ,
		PriPon:priPon,
		TwoStageSpectroscope : twostate ,
		TwoPon:twoPon,
		ONUAccount : data.ONUAccount,
		PonPort:data.PonPort,
	}
	return obj
}

func GroupTopString(str string) string {
	t := strings.LastIndex(str, "-")
	hash := str[0:t]
	return hash
}

func GroupTailString(str string) string {
	t := strings.LastIndex(str, "-")
	if len(str) >= t+1 {
		hash := str[(t+1):]
		return hash
	}
	return ""
}


