package routers

import (
	"pyg/pyg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister")
    //发送短信
    beego.Router("/sendMsg",&controllers.UserController{},"post:HandleSendMsg")
}
