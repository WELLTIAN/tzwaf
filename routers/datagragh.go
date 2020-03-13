package routers

import (
	//"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	//"strconv"
	//"strings"
	//"time"

	//"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"

	"github.com/WELLTIAN/tzwaf/models"
	"github.com/WELLTIAN/tzwaf/modules/util"
	//"github.com/WELLTIAN/tzwaf/setting"
)

func ListLogsApi(ctx *macaron.Context, sess session.Store) {
	log.Println("session,", sess.Get("uid"))
	if sess.Get("uid") != nil {
		log.Println(sess.Get("uid"))
		logs, _ := models.ListLog()
		//ctx.Data["Rules"] = logs
		//ctx.HTML(200, "logs")
		ctx.JSON(200, &logs)
	} else {
	    ret := util.Message{Status: 2, Message: "出错啦"}
		ctx.JSON(200, &ret)
	}
}

//攻击IP top5返回Json数据API 
func ListLogTopApi(ctx *macaron.Context, sess session.Store) {
	log.Println("session,", sess.Get("uid"))
	if sess.Get("uid") != nil {
		log.Println(sess.Get("uid"))
		logs, _ := models.ListLogTop()
		//ctx.Data["Rules"] = logs
		ctx.JSON(200, &logs)
	} else {
	    ret := util.Message{Status: 2, Message: "出错啦"}
		ctx.JSON(200, &ret)
	}
}

func ListLogTotalApi(ctx *macaron.Context, sess session.Store) {
	log.Println("session,", sess.Get("uid"))
	if sess.Get("uid") != nil {
		log.Println(sess.Get("uid"))
		logs, _ := models.ListLogTotal()
		//ctx.Data["Rules"] = logs
		ctx.JSON(200, &logs)
	} else {
	    ret := util.Message{Status: 2, Message: "出错啦"}
		ctx.JSON(200, &ret)
	}
}

func ListLogs(ctx *macaron.Context, sess session.Store) {
	log.Println("session,", sess.Get("uid"))
	if sess.Get("uid") != nil {
		log.Println(sess.Get("uid"))
		logs, _ := models.ListLog()
		ctx.Data["Logs"] = logs
		ctx.HTML(200, "Log")
	} else {
		ctx.Redirect("/login/")
	}
}