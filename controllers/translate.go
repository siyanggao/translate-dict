package controllers

import (
	"translate-dict/models"
	"translate-dict/services"

	"github.com/astaxie/beego"
)

type TransController struct {
	beego.Controller
}

func (t *TransController) Get() {
	word := t.GetString("word")
	res := new(models.BaseResponse)
	ok, data := services.Translate(word)
	if !ok {
		res.Code = 0
		res.Msg = "error"
	} else {
		res.Code = 1
		res.Data = data
	}
	t.Data["json"] = res
	t.ServeJSON()
}
