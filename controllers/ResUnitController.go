package controllers

import (
	"Dianshi/common"
	"Dianshi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Users
type ResUnitController struct {
	beego.Controller
}

// @router /getResUnit [post]
func (c *ResUnitController) GetResUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{common.FAIL, "参数错误", ""}
		c.ServeJSON()
		return
	}
	var resUnit models.TbResUnit
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &resUnit)
	db := models.Db
	fmt.Print(resUnit.ResId)
	qs := db.QueryTable("tb_res_unit")
	var resUnits []*models.TbResUnit
	_, err := qs.Filter("res_id", resUnit.ResId).All(&resUnits)
	if err != nil {
		c.Data["json"] = common.Result{common.FAIL, "查询错误", err.Error()}
	} else {
		//data,_ := json.Marshal(resUnits[0])
		c.Data["json"] = common.Result{common.SUCCESS, "", resUnits[0]}
	}
	c.ServeJSON()

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
//func (c *ResUnitController) GetAll() {
//	o := models.O
//	qs := o.QueryTable("tb_res_unit")
//	var resUnit []*models.TbResUnit
//	num, _ := qs.All(&resUnit)
//	fmt.Println(resUnit)
//	fmt.Println(num)
//	c.Data["json"] = resUnit
//	c.ServeJSON()
//}
