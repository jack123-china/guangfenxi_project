package common

type  Warning struct {
	Id      int         `json:"id" db:"id"`
	OLTName string   `db:"otlname"`//OLT名称
	Ponport string  `db:"ponport"`//"db:``"//错误端口
	PrimaryBeamSplitter  string   `db:"primarybeamsplitter"`
	TwoStageSpectroscope string   `db:"twostagespectroscope"`
	PriPon string 				    `db:"pripon"`
	TwoPon string 				    `db:"twopon"`
	Ip    string  	`db:"ip"`//错误ip
	ErrorType int  `db:"errortype"`//错误类型
	ErroeDiscrde int   `db:"errdiscribe"`      //错误描述
	DownMin  float64  `db:"downmin"`//最小下降幅度
	DownMax  float64   	`db:"downmax"`//最大下降幅度
	FaultNum int   `db:"falutnum"`//故障次数
	FaultFinalHappened string `db:"finalhappened"`//最后故障发生时间
	FaultHappenedTime int  `db:"heppenedtime"`//"db:`heppenedtime`"//故障持续时间   分钟
	FaultState int `db:"faultstate"`//"db:`faultstate`"//故障状态
	EffectNum int `db:"effectnum"`//"db:`effectnum`"//影响帐号数量
	ONUAccount string  `db:"account"`//"db:`onuaccount`"
	RecvOptPower string	 `db:"recvoptpower"`//"db:`recvoptpower`"//接收光功率
	TimeTips string `db:"tipstime"`
	IsSingleError int  `db:"isinglerror"` //是否包含在其他错误中  0 1
}

type Arv_time struct {
	Id      int         `db:"id"`
	ErrorType int `db:"errortype"`//错误类型
	ErroeDiscrde int   `db:"errdiscribe"`
	FaultNum int   `db:"falutnum"`//故障次数
}