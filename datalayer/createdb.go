package datalayer

import (
	"gfx_project/logsystem"
	"fmt"
)

//分光器是不变的 创建一次就好
func CreateSplliterTable() error {
	logsystem.Gfxlog.Info("======call CreateSplliterTable================\n")

	sql := `create table IF NOT EXISTS  splliterTable(
		ponport VARCHAR(100) NOT NULL,
		primarybeamsplitter VARCHAR(100) NOT NULL,
		twostagespectroscope VARCHAR(100) NOT NULL,
		pripon VARCHAR(10) NOT NULL,
		twopon VARCHAR(10) NOT NULL,
		onuaccount VARCHAR(50) NOT NULL
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	_ ,err  := DB.Exec(sql)
	if err != nil {
		logsystem.Gfxlog.Err("create table Splliter error = " , err)
		return err
	}
	sql = `alter table splliterTable add unique index ponwithaccount(ponport,onuaccount);`
	_ ,err  = DB.Exec(sql)
	if err != nil {
		logsystem.Gfxlog.Err("splliterTable add INDEX index_namee  error = " , err)
		return err
	}

	return  nil
}

//创建表
func CreateLogTable(name string) error {
	logsystem.Gfxlog.Info("======call CreateLogTable================\n")
	name = "tableLog" + name
	sql := `create table IF NOT EXISTS ` +   " %s " + `(
		id INT UNSIGNED AUTO_INCREMENT,
		time VARCHAR(20) NOT NULL default "",
		tipstime VARCHAR(20) NOT NULL default "",
		ip VARCHAR(20) NOT NULL,
		cityname VARCHAR(50) NOT NULL,
		otlname VARCHAR(100) NOT NULL,
		ponport VARCHAR(100) NOT NULL,
		primarybeamsplitter VARCHAR(100) NOT NULL,
		twostagespectroscope VARCHAR(100) NOT NULL,
		pripon VARCHAR(10) NOT NULL,
		twopon VARCHAR(10) NOT NULL,
		onuaccount VARCHAR(50) NOT NULL,
		onupasswd VARCHAR(30) NOT NULL,
		dispontoonu VARCHAR(30) NOT NULL,
		recvoptpower VARCHAR(50) NOT NULL,
		sendoptpower VARCHAR(50) NOT NULL,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	sqlstr := fmt.Sprintf(sql, name)
	//fmt.Printf("CreateLogTable sql 语句 = %v \n" , sqlstr)

	_ ,err  := DB.Exec(sqlstr)
	if err != nil {
		logsystem.Gfxlog.Err("create table error = " , err)
		return err
	}

	//index_onuaccount(onuaccount)
	slq := `alter table ` +   " %s " + ` add INDEX index_ponport(ponport) ;`
	sqlstr = fmt.Sprintf(slq,name)
	_ ,err  = DB.Exec(sqlstr)
	if err != nil {
		logsystem.Gfxlog.Err("add INDEX index_ponport  error = " , err)
		return err
	}

	slq = `alter table ` +   " %s " + ` add INDEX index_onuaccount(onuaccount) ;`
	sqlstr = fmt.Sprintf(slq,name)
	_ ,err  = DB.Exec(sqlstr)
	if err != nil {
		logsystem.Gfxlog.Err("add INDEX index_onuaccount  error = " , err)
		return err
	}
	return nil
}

//创建光距表
func CreateDisTable(name string)  error {
	logsystem.Gfxlog.Info("======call CreateDisTable================\n")
	name = "tableDis" + name
	sql := `create table IF NOT EXISTS ` +   " %s " + `(
		onupasswd VARCHAR(30) NOT NULL,
	 	onuupline VARCHAR(30) NOT NULL,
	 	onudownline VARCHAR(30) NOT NULL,
	 	dispontoonu VARCHAR(30) NOT NULL,
		UNIQUE(onupasswd)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	sqlstr := fmt.Sprintf(sql, name)
	_ ,err  := DB.Exec(sqlstr)
	if err != nil {
		logsystem.Gfxlog.Err("create table error = " , err)
		return err
	}
	return nil
}

//创建告警表
func CreateFaultTable(name string ) error{
	logsystem.Gfxlog.Info("======call CreateFaultTable================\n")
	//name := "falsealarm" + date

	slq := `
		create table IF NOT EXISTS ` + " %s " + `(
		id INT UNSIGNED AUTO_INCREMENT,
		otlname VARCHAR(50) NOT NULL default "",
		ponport VARCHAR(50) NOT NULL default "",
		tipstime VARCHAR(20) NOT NULL default "",
		primarybeamsplitter VARCHAR(100) NOT NULL,
		twostagespectroscope VARCHAR(100) NOT NULL,
		pripon VARCHAR(10) NOT NULL,
		twopon VARCHAR(10) NOT NULL,
		ip  VARCHAR(20) NOT NULL default "",
		errortype int unsigned not null default 0,
		errdiscribe int unsigned not null default 0,
		downmin decimal(12, 10) not null default 0 ,
		downmax decimal(12, 10) not null default 0 ,
		falutnum int unsigned not null default 0,
		finalhappened VARCHAR(20) NOT NULL default "",
		heppenedtime int  not null default 0,
		faultstate tinyint not null default 0,
		effectnum int unsigned not null default 0,
		account VARCHAR(20) NOT NULL default "",
		recvoptpower VARCHAR(20) NOT NULL,
		isinglerror  int  not null default 0,
		PRIMARY KEY ( id )
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`
	slq = fmt.Sprintf(slq , name)
	_ ,err  := DB.Exec(slq)
	if err != nil {
		logsystem.Gfxlog.Err("CreateFaultTable table error = " , err)
		return err
	}
	return nil
}

func CreateUserGroupTable() (err error){
	slq:= `
		create table IF NOT EXISTS  only_usergroup(
			id INT UNSIGNED AUTO_INCREMENT,
			otlname VARCHAR(50) NOT NULL default "",
			address VARCHAR(20) NOT NULL,
			ponport VARCHAR(100) NOT NULL,
			primarybeamsplitter VARCHAR(100) NOT NULL,
			twostagespectroscope VARCHAR(100) NOT NULL,
			pripon VARCHAR(10) NOT NULL,
			twopon VARCHAR(10) NOT NULL,
			onuaccount VARCHAR(50) NOT NULL,
			PRIMARY KEY ( id )
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`


	_ ,err = DB.Exec(slq)
	if err != nil {
		logsystem.Gfxlog.Err("CreateUserGroupTable table error = " , err)
		return err
	}


	slq = `ALTER TABLE only_usergroup ADD INDEX ponindex(ponport );`
	_ ,err  = DB.Exec(slq)
	if err != nil {
		logsystem.Gfxlog.Err("add INDEX ponindex  error = " , err)
		return err
	}

	slq = `ALTER TABLE only_usergroup ADD INDEX accpount(onuaccount );`
	_ ,err  = DB.Exec(slq)
	if err != nil {
		logsystem.Gfxlog.Err("add INDEX accpount  error = " , err)
		return err
	}
	return
}

func CreateAvrTime(name string ) (err error){
	logsystem.Gfxlog.Info("======call CreateAvrTime================\n")

	slq := `
		create table IF NOT EXISTS ` + " %s " + `(
		id INT UNSIGNED AUTO_INCREMENT,
		errortype int unsigned not null default 0,
		errdiscribe int unsigned not null default 0,
		falutnum int unsigned not null default 0,
		PRIMARY KEY ( id )
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`
	slq = fmt.Sprintf(slq , name)
	_ ,err = DB.Exec(slq)
	if err != nil {
		logsystem.Gfxlog.Err("CreateFaultTable table error = " , err)
		return err
	}
	return nil
}
