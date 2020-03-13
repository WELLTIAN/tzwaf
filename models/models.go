package models

import (
	"fmt"
	"log"

	"github.com/WELLTIAN/tzwaf/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	err    error
)

func init() {
	sec := setting.Cfg.Section("database")
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		sec.Key("USER").String(),
		sec.Key("PASSWD").String(),
		sec.Key("HOST").String(),
		sec.Key("NAME").String()))
	if err != nil {
		log.Panicf("Faild to connect to database, err:%v", err)
	}

	// log.Println(Engine, err)

	//使用xorm创建表，这儿我们创建四张表，分别为site user rules logs
	Engine.Sync2(new(Site))
	Engine.Sync2(new(User))
	Engine.Sync2(new(Rules))
	Engine.Sync2(new(Logs))

	ret, err := Engine.IsTableEmpty(new(User))
	if err == nil && ret {
		log.Printf("create new user:%v, password:%v\n", "admin", "123456")
		NewUser("admin", "123456")
	}

	ret, err = Engine.IsTableEmpty(new(Rules))
	if err == nil && ret {
		log.Println("Insert default waf rules")
		Engine.Exec(DefaultRules)
	}
}
