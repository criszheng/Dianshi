package controllers

import (
	"Dianshi/common"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"log"
	"path"
)

// Operations about Users
type PicController struct {
	beego.Controller
}

// @router /savePicture [post]
func (c *PicController) SavePicture() {
	files, err := c.GetFiles("files")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	for _, file := range files {
		ext := path.Ext(file.Filename)
		//验证后缀名是否符合要求
		var AllowExtMap map[string]bool = map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}
		if _, ok := AllowExtMap[ext]; !ok {
			c.Data["json"] = common.Result{Code: common.FAIL, Message: "后缀名不符合上传要求", Data: ""}
			c.ServeJSON()
			return
		}
	}
	//上传对象
	cosUtil := common.CosUtil{}
	cos := cosUtil.CreateCosClient()
	var urls []string
	for _, file := range files {
		uid, _ := uuid.NewV4()
		filePath := "pic/" + uid.String() + path.Ext(file.Filename)
		data, _ := file.Open()
		common.PutFile(cos, filePath, data)
		defer data.Close()
		url := common.COSURL + "/" + filePath
		urls = append(urls, url)
	}
	c.Data["json"] = common.Result{Code: common.SUCCESS, Data: urls}
	c.ServeJSON()
}
