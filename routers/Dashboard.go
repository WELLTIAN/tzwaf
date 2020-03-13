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

	//"github.com/WELLTIAN/tzwaf/models"
	//"github.com/WELLTIAN/tzwaf/modules/util"
	//"github.com/WELLTIAN/tzwaf/setting"
)

func Dashboardhtml(ctx *macaron.Context, sess session.Store) {
	log.Println("session,", sess.Get("uid"))
	if sess.Get("uid") != nil {
		log.Println(sess.Get("uid"))
		ctx.HTML(200, "Dashboard")
	} else {
		ctx.Redirect("/login/")
	}
}