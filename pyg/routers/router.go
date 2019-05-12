package routers

import (
	"pyg/pyg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //发送短信
    beego.Router("/sendMsg",&controllers.UserController{},"post:HandleSendMsg")
    //邮箱激活
    beego.Router("/register-email",&controllers.UserController{},"get:ShowEmail;post:HandleEmail")
    //激活用户
    beego.Router("/active",&controllers.UserController{},"get:Active")
    //登录实现
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    //首页展示
    beego.Router("/article/index",&controllers.UserController{},"get:ShoeIndex;post:ShowIndex")
}
