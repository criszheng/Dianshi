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

// @router /addResUnit [post]
func (c *ResUnitController) AddResUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var resUnit models.TbResUnit
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &resUnit)
	db := models.Db
	fmt.Println(resUnit.ResName)
	id, err := db.Insert(&resUnit)
	beego.Notice("AddResUnit, id is " + string(id))
	if err != nil {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "AddResUnit插入错误", Data: err.Error()}
	} else {
		resUnit.ResId = int(id)
		c.Data["json"] = common.Result{Code: common.SUCCESS, Data: resUnit}
	}
	c.ServeJSON()
}

// @router /getAllResUnit [post]
func (c *ResUnitController) GetAllResUnit() {
	db := models.Db
	qs := db.QueryTable("tb_res_unit")
	var resUnits []*models.TbResUnit
	_, err := qs.All(&resUnits)
	if err != nil {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "GetAllResUnit查询错误", Data: err.Error()}
	} else {
		for _, resUnit := range resUnits {
			qs := db.QueryTable("tb_ad_unit")
			count, _ := qs.Filter("res_id", resUnit.ResId).Count()
			resUnit.Count = int(count)
		}
		c.Data["json"] = common.Result{Code: common.SUCCESS, Data: resUnits}
	}
	c.ServeJSON()
}

// @router /getResUnit [post]
func (c *ResUnitController) GetResUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var resUnit models.TbResUnit
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &resUnit)
	db := models.Db
	qs := db.QueryTable("tb_res_unit")
	var resUnits []*models.TbResUnit
	_, err := qs.Filter("res_id", resUnit.ResId).All(&resUnits)
	if err != nil {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "GetResUnit查询错误", Data: err.Error()}
	} else {
		//data,_ := json.Marshal(resUnits[0])
		c.Data["json"] = common.Result{Code: common.SUCCESS, Data: resUnits[0]}
	}
	c.ServeJSON()

}

// @router /updateResUnit [post]
func (c *ResUnitController) UpdateResUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var resUnit models.TbResUnit
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &resUnit)
	db := models.Db
	num, err := db.Update(&resUnit)
	beego.Notice("UpdateResUnit, num is " + string(num))
	if err != nil {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "UpdateResUnit更新错误", Data: err.Error()}
	} else {
		c.Data["json"] = common.Result{Code: common.SUCCESS, Data: num}
	}
	c.ServeJSON()
}

// @router /deleteResUnit [post]
func (c *ResUnitController) DeleteResUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var resUnit models.TbResUnit
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &resUnit)
	db := models.Db
	num, err := db.Delete(&resUnit)
	beego.Notice("UpdateResUnit, num is " + string(num))
	if err != nil {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "DeleteResUnit删除错误", Data: err.Error()}
	} else {
		c.Data["json"] = common.Result{Code: common.SUCCESS, Data: num}
	}
	c.ServeJSON()
}
