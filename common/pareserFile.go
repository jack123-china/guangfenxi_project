package common

//数据采集结构体

//分光器
type Splitter struct {
	//PON端口，一级分光器，二级分光器，
	PonPort string                `db:"ponport"`
	PrimaryBeamSplitter  string   `db:"primarybeamsplitter"`
	PriPon string 				  `db:"pripon"`
	TwoStageSpectroscope string   `db:"twostagespectroscope"`
	TwoPon string 				  `db:"twopon"`
	ONUAccount string             `db:"onuaccount"`
}

//光功率
type RawLog struct {
	PonPort  string //OLT PON端口
	Time string//时间
	CityName string//城市
	OLTName string//OLT名称
	OLTMgrIp string//OLT管理IP
	OUNumber  string//ONU编号
	ONUAccount string//ONU账号
	ONUPassWd string//ONU密码
	RecvOptPower string//接收光功率
	SendOptPower string//发光功率
	CPUOccupy int//CPU占用
	MemOccupy int//内存占用
	GateWayRunTime string//网关运行时长(秒)
	TipsTime  string    //时间标志
}

//光距
type OptDistance struct {
	ONUPassWd string             `db:"onupasswd"`
	ONUUpLine string  			 `db:"onuupline"`
	ONUDownLine  string  		 `db:"onudownline"`
	DistanceFromPONToONU string  `db:"dispontoonu"`
}

//存数据库的表
type CompositeData struct {
	Id int `db:"id"`
	CityName string     			`db:"cityname"`//城市
	Ip string                       `db:"ip"`
	OLTName string          		`db:"otlname"`//OLT名称
	PonPort  string 				`db:"ponport"`//OLT PON端口
	PrimaryBeamSplitter  string   	`db:"primarybeamsplitter"`
	PriPon string 				    `db:"pripon"`
	TwoStageSpectroscope string   	`db:"twostagespectroscope"`
	TwoPon string 				    `db:"twopon"`
	ONUAccount string				`db:"onuaccount"`//ONU账号
	ONUPassWd string				`db:"onupasswd"`//ONU密码
	RecvOptPower string			 	`db:"recvoptpower"`//接收光功率
	SendOptPower string			 	`db:"sendoptpower"`//发光功率
	DistanceFromPONToONU string  	`db:"dispontoonu"`	//光距
	FinalTime string                `db:"time"` //最后发生的时间
	TipsTime  string    			`db:"tipstime"`//时间标志
}

type ParseFileData struct {
	Mp_log_spliter map[string]* RawLog

	Mp_sp map[string]* Splitter
	Mp_dis  map[string]* OptDistance
	FinalData []*CompositeData
}

type UserGroup struct {
	Id int `db:"id"`
	PonPort  string `db:"ponport"`//OLT PON端口
	Address string `db:"address"`
	OLTName string `db:"otlname"`
	PrimaryBeamSplitter  string   	`db:"primarybeamsplitter"`
	PriPon string 				    `db:"pripon"`
	TwoStageSpectroscope string   	`db:"twostagespectroscope"`
	TwoPon string 				    `db:"twopon"`
	ONUAccount string				`db:"onuaccount"`//ONU账号
}

