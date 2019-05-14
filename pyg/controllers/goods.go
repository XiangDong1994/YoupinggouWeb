package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pyg/pyg/models"
)

type GoodsController struct {
	beego.Controller
}

func(this*GoodsController)ShowIndex(){
	name := this.GetSession("name")
	if name != nil {
		this.Data["name"] = name.(string)
	}else {
		this.Data["name"] = ""
	}

	//获取类型信息并传递给前段
	//获取一级菜单
	o := orm.NewOrm()
	//接受对象
	var oneClass []models.TpshopCategory
	//查询
	o.QueryTable("TpshopCategory").Filter("Pid",0).All(&oneClass)


	//获取第二级
	var types []map[string]interface{}
	for _,v := range oneClass{
		//行容器
		t := make(map[string]interface{})

		var temp []models.TpshopCategory
		o.QueryTable("TpshopCategory").Filter("Pid",v.Id).All(&temp)
		t["t1"] = v
		t["t2"] = temp
		types = append(types,t)
	}

	//获取三级菜单
	for _,v1:=range types{
		//循环得到了二级菜单
		for _,v2:=range v1{
			t:=make(map[string]interface{})
			var thirdClass []models.TpshopCategory
		//	获取三级菜单
		o.QueryTable("TpshopCategory")
		}
	}

	this.Data["types"] = types
	this.TplName = "index.html"
}
