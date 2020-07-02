package warnning

//故障类型
const (
	FIRST_LEVEL_LINK = 100  //一级链路
	FIRST_SPLLITER = 101    //一级分光器
	SECOND_LEVEL_LINK = 102 //二级链路
	SECOND_SPLLITER = 103   //二级分光器
	PIGTAIL_OR_CONN = 104   //尾纤或连接器
)


const (
	Child_Deter_Light_Path = 200  //光路劣化
	Child_Optical_Interruption = 201//光路中断
	Child_FIRST_LEVEL_LINK = 202//一级链路
	Child_SECOND_LEVEL_LINK = 203//二级链路

	Child_Weak_light = 204   //弱光
	Child_Extremely_Weak_light = 205  //极度弱光
	Child_Strong_Light = 206//强光

	CHild_ONU_OUTLINE = 207 //ONU用户下线
)

//收发光功率  上限 下限的阀值
const (
	Stand_Recv_OutLine_1 = 0 //0  -40 是用户下线
	Stand_Recv_OutLine_2 = -40

	Standard_Recv_Raw_up = -8
	Standard_Recv_Raw_dpwn = -27
	Standard_Recv_Extremely_Weak_light = -34  //极度弱光的阀值


	Standard_Send_Raw_up = 7
	Standard_Send_Raw_dpwn = -1
)