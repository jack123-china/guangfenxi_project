package datacollection

import (
	"fmt"
	"gfx_project/common"
	db "gfx_project/datalayer"
	"gfx_project/logsystem"
	"gfx_project/util"
	"sync"
)

var OrigilData *common.ParseFileData
var CurrentDate string
var yesterdayDisDate string  //目前光距的表

func init() {
	OrigilData = &common.ParseFileData{
		Mp_log_spliter  : make(map[string]* common.RawLog),
		Mp_sp : make(map[string]* common.Splitter),
		Mp_dis :make(map[string]* common.OptDistance),
		FinalData :make([]*common.CompositeData,0),
	}

	CurrentDate = util.GetToday()
	//fmt.Printf("当前日期 = %v \n" , CurrentDate)
	yesterdayDisDate = util.GetYesterday()
}

func InitAllData(logpath string, spliterpath string , disPath string) {
	logsystem.Gfxlog.Info("call InitAllData ")
	//add := &common.ParseFileData{}
	go db.CreateUserGroupTable()

	wg := sync.WaitGroup{}
	wg.Add(2)

	//go goParserLogFile(&wg , logpath , OrigilData)
	go goParserSpliterFile(&wg , spliterpath , OrigilData)
	go goParserDisFile(&wg , disPath , OrigilData)

	wg.Wait()

	fmt.Println("完成解析\n")

	if db.IsExitsTable("splliterTable" , "ponport") == false {
		err := db.CreateSplliterTable()
		if err != nil {
			logsystem.Gfxlog.Info("IsExistsSplliterTables CreateSplliterTable create table error")
			return
		}
	}

	go db.PreInsertSplliter(OrigilData.Mp_sp)

	db.CreateDisTable(yesterdayDisDate)

	//db.CreateAvrTime("arvtimetable")

	go db.PreInsertDis(OrigilData.Mp_dis , yesterdayDisDate)
	fmt.Println("分光器信息插入数据库完成\n")
}

func goParserLogFile(w *sync.WaitGroup ,path string, add *common.ParseFileData)  {
	add.Mp_log_spliter =  PareseRawLogFile(path)
	w.Done()
}

func goParserSpliterFile(w *sync.WaitGroup ,path string , add *common.ParseFileData)  {
	add.Mp_sp = PareseSplliter(path)
	w.Done()
}

func goParserDisFile(w *sync.WaitGroup ,path string , add *common.ParseFileData)  {
	add.Mp_dis = ParserDistance(path)
	w.Done()
}
