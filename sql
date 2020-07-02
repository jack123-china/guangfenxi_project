/*

//创建 分光器表
此处ponport + onuaccount 字段必须为唯一
这地方需要加id 递增作为主键吗？？？
create table IF NOT EXISTS  splliterTable(
     `ponport` VARCHAR(100) NOT NULL,
	 `primarybeamsplitter` VARCHAR(100) NOT NULL,
	 `twostagespectroscope` VARCHAR(100) NOT NULL,
	 `onuaccount` VARCHAR(50) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table splliterTable add unique index ponwithaccount(ponport,onuaccount);


create table IF NOT EXISTS  distancetable(
   `onupasswd` VARCHAR(30) NOT NULL,
	 `onuupline` VARCHAR(30) NOT NULL,
	 `onudownline` VARCHAR(30) NOT NULL,
	 `dispontoonu` VARCHAR(30) NOT NULL,
		UNIQUE(onupasswd)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table IF NOT EXISTS  table20200203(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `cityname` VARCHAR(50) NOT NULL,
    ip VARCHAR(20) NOT NULL,
    `otlname` VARCHAR(100) NOT NULL,
	 `ponport` VARCHAR(100) NOT NULL,
	  `primarybeamsplitter` VARCHAR(100) NOT NULL,
	 `twostagespectroscope` VARCHAR(100) NOT NULL,
	 `onuaccount` VARCHAR(50) NOT NULL,
	 `onupasswd` VARCHAR(30) NOT NULL,
	 `dispontoonu` VARCHAR(30) NOT NULL,
	 `recvoptpower` VARCHAR(50) NOT NULL,
	 `sendoptpower` VARCHAR(50) NOT NULL,
	  PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE table20200203 ADD INDEX index_name (ponport , onuaccount)

create table IF NOT EXISTS falsealarm (
    `id` INT UNSIGNED AUTO_INCREMENT,
    `otlname` VARCHAR(50) NOT NULL default "",
    `ponport` VARCHAR(50) NOT NULL default "",
    `ip`  VARCHAR(20) NOT NULL default "",
    `errortype` int unsigned not null default 0,
    `errdiscribe` int unsigned not null default 0,
    `downmin` decimal(12, 10) not null default 0 ,
     `downmax` decimal(12, 10) not null default 0 ,
     `falutnum` int unsigned not null default 0,
     `finalhappened` timestamp NOT NULL,
     `heppenedtime` decimal(5, 2)  not null default 0,
     `faultstate` tinyint not null default 0,
      `effectnum` int unsigned not null default 0,
      `ishow` tinyint not null default 0,
        PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table IF NOT EXISTS  only_usergroup(
    id INT UNSIGNED AUTO_INCREMENT,
    otlname VARCHAR(50) NOT NULL default "",
    address VARCHAR(20) NOT NULL,
	ponport VARCHAR(100) NOT NULL,
	primarybeamsplitter VARCHAR(100) NOT NULL,
	twostagespectroscope VARCHAR(100) NOT NULL,
	`onuaccount` VARCHAR(50) NOT NULL,
	PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE table20200203 ADD INDEX ponindex(ponport )
ALTER TABLE table20200203 ADD INDEX accpount(onuaccount )
*/