package routers

import (
	"github.com/astaxie/beego"
	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/controllers"
)

func Init() {
	beego.Router("/hello_world", &controllers.HelloController{})
	beego.Router("/create_user", &controllers.UserController{}, "get,post:CreateUser")
	beego.Router("/users", &controllers.UserController{}, "get:ListUsers")
}
