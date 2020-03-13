package models

import (
	"log"
)

type Logs struct {
	Id          int64
	User_agent  string `xorm:"varchar(200) not null"`
	Attack_type string `xorm:"varchar(100) not null"`
	Req_url     string `xorm:"varchar(100) not null"`
	Client_ip   string `xorm:"varchar(100) not null"`
	Local_time  string `xorm:"varchar(100) not null"`
	Rule_tag    string `xorm:"varchar(100) not null"`
	Req_data    string `xorm:"varchar(100) not null"`
	Server_name string `xorm:"varchar(100) not null"`
	Count int64   //这儿的Count值作为查询Top的辅助使用，值为Null	
}

func ListLog() (logs []Logs, err error) {
	logs = make([]Logs, 0)
	err = Engine.Desc("Id").Find(&logs)
	log.Println(err, logs)
	return logs, err
}

//返回攻击IP top 5
func ListLogTop()(logs []Logs, err error) { 
	logs = make([]Logs, 0)
    err  = Engine.Sql("select client_ip,count(*) as Count from logs group by client_ip desc limit 5").Find(&logs)
	log.Println(err, logs)
	return logs, err
}

//返回总攻击条数
func ListLogTotal()(logs []Logs, err error) { 
	logs = make([]Logs, 0)
    err  = Engine.Sql("select count(*) as Count from logs").Find(&logs)
	log.Println(err, logs)
	return logs, err
}

//返回攻击类型top5
func ListLogAttaktypeTop()(logs []Logs, err error) { 
	logs = make([]Logs, 0)
    err  = Engine.Sql("select Attack_type,count(*) as Count from logs group by Attack_type desc limit 5").Find(&logs)
	log.Println(err, logs)
	return logs, err
}



