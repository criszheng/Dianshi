package controllers

import (
	"Dianshi/common"
	"Dianshi/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Users
type AdUnitController struct {
	beego.Controller
}

const (
	MOVE_TOP = 0
	MOVE_UP = 1
	MOVE_DOWN = 2
)

type UpdateOrderAttr struct{
   MoveOperation int
   AdId int
}

// @router /clickAd [post]
func (c *AdUnitController) ClickAd() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var adUnit models.TbAdUnit
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&adUnit); err == nil {    //传来addId

		db:=models.Db
		adUnit.ViewNum += 1
		num, err := db.Update(&adUnit)
		if err != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "增加点击量错误", Data: err.Error()}
		} else {
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: num}
		}
	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data:err.Error()}
	}

	c.ServeJSON()
}


// @router /updateOrder [post]
func (c *AdUnitController) UpdateOrder() {
	if len(c.Ctx.Input.RequestBody) == 0 {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "参数错误", Data: ""}
		c.ServeJSON()
		return
	}
	var updateOrderAttr UpdateOrderAttr
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &updateOrderAttr); err == nil {
		db := models.Db
		var adUnitChoosed models.TbAdUnit
		qs := db.QueryTable("tb_ad_unit")
		err1 :=db.QueryTable("tb_ad_unit").Filter("ad_id",updateOrderAttr.AdId).One(&adUnitChoosed)
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "UpdateOrder错误", Data: err1.Error()}
			return
		}
		var adUnits []*models.TbAdUnit
		_, err1 = qs.All(&adUnits)
		if err1 != nil {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "UpdateOrder错误", Data: err1.Error()}
		} else {
			c.Data["json"] = common.Result{Code: common.SUCCESS, Data: adUnits}
		}
		switch updateOrderAttr.MoveOperation {
		case MOVE_TOP:
			for _,adUnit := range adUnits{
				//从0还是1开始排
				if adUnit.AdID > adUnitChoosed.Order{
					continue
				} else if adUnit.AdID == updateOrderAttr.AdId {
					adUnit.Order = 0
				}else if adUnit.Order < adUnitChoosed.Order{
					adUnit.Order += 1
				}
				if _,err1 = db.Update(&adUnit);err1==nil{
				}else{
					c.Data["json"] = common.Result{Code: common.FAIL, Message: "更新数据库信息错误", Data: err1.Error()}
				}
			}
		case MOVE_UP:  //第一个还能上移吗，前端应该能识别并阻止此操作，不会往后端发请求吧
			for _,adUnit := range adUnits{
				if adUnit.Order == adUnitChoosed.Order {
					adUnit.Order -= 1
				}else if adUnit.Order == adUnitChoosed.Order - 1{
					adUnit.Order += 1
				}else{
					continue
				}
				if _,err1 = db.Update(&adUnit);err1==nil{
				}else{
					c.Data["json"] = common.Result{Code: common.FAIL, Message: "更新数据库信息错误", Data: err1.Error()}
				}
			}
		case MOVE_DOWN:
			for _,adUnit := range adUnits{
				if adUnit.Order == adUnitChoosed.Order {
					adUnit.Order += 1
				}else if adUnit.Order == adUnitChoosed.Order + 1{
					adUnit.Order -= 1
				}else{
					continue
				}
				if _,err1 = db.Update(&adUnit);err1==nil{
				}else{
					c.Data["json"] = common.Result{Code: common.FAIL, Message: "更新数据库信息错误", Data: err1.Error()}
				}
			}
		}

	} else {
		c.Data["json"] = common.Result{Code: common.FAIL, Message: "传入参数错误", Data:err.Error()}
	}
	c.ServeJSON()
}





