package datacollection

import (
	"fmt"
	"gfx_project/common"
	db "gfx_project/datalayer"
	"gfx_project/logsystem"
	"gfx_project/util"
	"gfx_project/warnning"
	"strconv"
	"strings"
)

func AddData(filetype int ,name string) {
	path := name
	log := make(map[string]*common.RawLog)

	if filetype == 0 {
		log  = PareseRawLogFile(path)
		name1 := strings.Split(name , "/")    //这地方是根据命名区出日期
		tmp := strings.Split(name1[len(name1) - 1] , ".")

		date := tmp[0][(len(tmp[0]) - 8) : ]

		time1, err  := strconv.Atoi(date)

		if err != nil  {
			logsystem.Gfxlog.Err("日期time1出现错误格式",date )
			return
		}
		time2 , err := strconv.Atoi(CurrentDate)
		if err != nil  {
			logsystem.Gfxlog.Err("日期time2出现错误格式",date )
			return
		}

		if time1 < time2 {
			logsystem.Gfxlog.Err("日期小于当日日期 ，格式不正确",time1 , time2 )
			return
		}

		if db.IsExitsTable("tableLog" + CurrentDate , "ponport")  == false {
			db.CreateLogTable(date)
			CurrentDate = date

			minTime := getMinTime(log)
			util.SetTimeTips(minTime)
			for _, v := range log {
				v.TipsTime = minTime
			}
		}else {
			nextTime := util.AfterFiveMin(db.GetMaxTimeTips("tableLog" + CurrentDate))//插入数据库之前就对时间作出修改
			util.SetTimeTips(nextTime)
			for _, v := range log {
				v.TipsTime = nextTime
			}

		}

		mc := CompareAddData(log)
		go warnning.RunPlice(mc , CurrentDate)

	}else if filetype == 1 { //分光器
		mp_sp := PareseSplliter(path)
		for k , v := range  mp_sp {
			OrigilData.Mp_sp[k] = v
		}
	}else { //光距
		dis :=ParserDistance(path)
		for k , v := range dis {
			OrigilData.Mp_dis[k] = v
		}

		//fmt.Printf("解析出来的光距数据 = %v \n" , OrigilData.Mp_dis)
		logDay := util.GetYesterday()
		yesterdayDisDate = logDay
		db.CreateDisTable(yesterdayDisDate)

		go db.ChangeFinalTableDis(yesterdayDisDate , OrigilData.Mp_dis)

		go db.PreInsertDis(OrigilData.Mp_dis, yesterdayDisDate)
	}
}

//组合
func CompareAddData(log map[string]*common.RawLog) []*common.CompositeData{
	mp := make([]*common.CompositeData , 0)
	usr := make([]*common.UserGroup , 0)
	fmt.Printf("当前数据长素的    =====   %d\n" , len(log))
	var time string
	var timeTips string
	for _, val := range log {
		sp := db.GetSpliterData(val.PonPort, val.ONUAccount)
		if len(sp) != 1 {
			logsystem.Gfxlog.Err("找到的分光器信息不是一条。。。",len(sp) )
			continue
		}

		dis := db.GetDistanceData(val.ONUPassWd ,yesterdayDisDate)
		if len(dis) == 0 {
			logsystem.Gfxlog.Err("找到的光距信息不是一条")
			continue
		}

		data := NewCompositeData(val, sp[0] , dis[0])
		mp = append(mp, data)
		//db.SetParserDataToDB(data , currentDate)

		tmp :=  NewUserGroup(val , data.PrimaryBeamSplitter,data.TwoStageSpectroscope , data.PriPon, data.TwoPon)
		usr = append(usr ,tmp)

		time = val.Time
		timeTips  = val.TipsTime
	}
	dbUser := db.GetAllUserData()
	temp := CheckLostUserLog(usr , dbUser ,time , timeTips )
	fmt.Printf("当前的错误新增 = %v \n" , len(temp))

	mp = append(mp , temp...)
	go db.CheckUsergroup(usr)
	go InsertCompareData(mp , CurrentDate)

	return mp
}

//新增没有用户日志的
func CheckLostUserLog(logUser []*common.UserGroup,dbUser []*common.UserGroup , time string , tips string )  []*common.CompositeData{
	mp := make(map[string] *common.UserGroup)

	log := make([]*common.CompositeData ,0)
	for _ , v := range  logUser {
		mp[v.ONUAccount] = v
	}
	fmt.Printf("GetLogDataByUser s数据长度 = %v  ==== %v \n" , len(logUser) , len(dbUser))
	for _ ,v := range  dbUser {
		if _, ok := mp[v.ONUAccount] ; !ok {
			tmp ,err  := db.GetLogDataByUser(v)
			fmt.Println("===============CheckLostUserLog======================\n")
			if err != nil || len(tmp) == 0 {
				fmt.Printf("GetLogDataByUser error = %v \n" , err)
				continue
			}
			temp := NewConPositeLog(tmp[0] ,time,  tips)
			log = append(log , temp)
		}
	}
	return log
}

func InsertCompareData(data []*common.CompositeData,date string) {
	logsystem.Gfxlog.Err("call InsertCompareData ", len(data))

	if !check(data) {
		str := db.GetMaxTimeTips("tableLog" + date)
		if str == "" {
			logsystem.Gfxlog.Err("时间标志错误,不能入库")
			return
		}
		for _, v := range  data {
			v.TipsTime = str
			break
		}
	}

	for _, v := range  data {
		if v.RecvOptPower == "" {
			v.RecvOptPower = "0"
		}

		if v.SendOptPower == "" {
			v.SendOptPower = "0"
		}
		db.SetParserDataToDB(v , date)
	}
}


func check(data []*common.CompositeData) bool {
	if len(data) == 0 {
		return false
	}

	fmt.Printf("检查之后的str = %v \n" , data[0].TipsTime)
	return data[0].TipsTime != ""
}

//获得最小的时间
func getMinTime(data map[string]*common.RawLog) string{
	var min string
	min = "2222-03-21 09:04:25"
	for _ , v := range  data {
		ok := util.CompareTime(min, v.TipsTime)
		if !ok {
			min = v.TipsTime
		}
	}

	//fmt.Printf("获得最小的时间 = %v  ,当前日志的长度 = %v \n" ,min ,len(data) )
	return min
}