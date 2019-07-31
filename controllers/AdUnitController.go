package controllers

import (
	"Dianshi/common"
	"Dianshi/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type AdUnitController struct {
	beego.Controller
}

//新建广告位
// @router /addAdUnit [post]
func (c *AdUnitController) AddAdUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var adUnit models.TbAdUnit
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &adUnit); err == nil {
		db := models.Db
		id, err1 := db.Insert(&adUnit)
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "AddResUnit插入错误", Data: err1.Error()}
		} else {
			adUnit.AdID = int(id)
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: adUnit}
		}
	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data:err.Error()}
	}
	c.ServeJSON()
}

//根据resId查询广告位列表
// @router /getAllAdUnitByRes [post]
func (c *AdUnitController) GetAllAdUnitByRes() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var adUnit models.TbAdUnit
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &adUnit); err == nil {
		db := models.Db
		qs := db.QueryTable("tb_ad_unit")
		var adUnits []*models.TbAdUnit
		_, err1 := qs.Filter("res_id", adUnit.ResId).All(&adUnits)
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "GetAllAdUnitByRes查询错误", Data: err1.Error()}
		} else {
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: adUnits}
		}
	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data:err.Error()}
	}
	c.ServeJSON()
}

// @router /updateAdUnit [post]
func (c *AdUnitController) UpdateAdUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var adUnit models.TbAdUnit
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &adUnit); err == nil {
		db := models.Db
		num, err1 := db.Update(&adUnit)
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "UpdateAdUnit更新错误", Data: err1.Error()}
		} else {
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: num}
		}
	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data:err.Error()}
	}
	c.ServeJSON()
}

// @router /deleteAdUnit [post]
func (c *AdUnitController) DeleteAdUnit() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var adUnit models.TbAdUnit
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &adUnit); adUnit.AdID != 0 && err == nil {
		db := models.Db
		num, err1 := db.Delete(&models.TbAdUnit{AdID : adUnit.AdID})
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "UpdateAdUnit删除错误", Data: err1.Error()}
		} else {
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: num}
		}
	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data: err.Error()}
	}
	c.ServeJSON()
}