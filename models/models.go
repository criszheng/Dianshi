package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//ormer
var Db orm.Ormer

type TbUser struct {
	Uid        int       `orm:"column(uid);pk" json:"uid"` // 设置主键
	RtxName    string    `json:"rtxName"`
	Password   string    `json:"password"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
}

type TbAdUnit struct {
	AdId          int       `orm:"column(ad_id);pk" json:"adId"` // 设置主键
	AdName        string    `json:"adName"`
	AdSubtitle    string    `json:"adSubtitle"`
	ResId         int       `json:"resId"`
	Type          int       `json:"type"`
	Order         int       `json:"order"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`
	CreatorUid    string    `json:"creatorUid"`
	UpdateUid     string    `json:"updateUid"`
	RelevantUsers string    `json:"relevantUsers"`
	IsPublished   int       `json:"isPublished"`
	Title         string    `json:"title"`
	BgType        int       `json:"bgType"`
	BgColor       string    `json:"bgColor"`
	BgPicUrl      string    `json:"bgPicUrl"`
	LinkType      int       `json:"linkType"`
	ButtonName    string    `json:"buttonName"`
	Link          string    `json:"link"`
	OnTime        time.Time `json:"onTime"`
	OffTime       time.Time `json:"offTime"`
	Remark        string    `json:"remark"`
	IsUrgency     int       `json:"isUrgency"`
	UrgencyMsg    string    `json:"urgencyMsg"`
	IsUrgencyLink int       `json:"isUrgencyLink"`
	UrgencyUrl    string    `json:"urgencyUrl"`
	ViewNum       int       `json:"viewNum"`
}

type TbResUnit struct {
	ResId         int       `orm:"column(res_id);pk" json:"resId"` // 设置主键
	ResName       string    `json:"resName"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`
	CreatorUid    string    `json:"creatorUid"`
	UpdateUid     string    `json:"updateUid"`
	RelevantUsers string    `json:"relevantUsers"`
	IsPublished   int       `json:"isPublished"`
	Count         int       `orm:"-" json:"count"`
	ResTitle      string    `json:"resTitle"`
	LinkType      int       `json:"linkType"`
	ButtonName    string    `json:"buttonName"`
	Link          string    `json:"link"`
	Url           string    `json:"url"`
	BgType        int       `json:"bgType"`
	BgColor       string    `json:"bgColor"`
	BgPicUrl      string    `json:"bgPicUrl"`
	Msg           string    `json:"msg"`
	Remark        string    `json:"remark"`
	IsUrgency     int       `json:"isUrgency"`
	UrgencyMsg    string    `json:"urgencyMsg"`
	IsUrgencyLink int       `json:"isUrgencyLink"`
	UrgencyUrl    string    `json:"urgencyUrl"`
}

type TbPic struct {
	PicId      int       `orm:"column(pic_id);pk" json:"picId"` // 设置主键
	AdId       int       `json:"adId"`
	Url        string    `json:"url"`
	Link       string    `json:"link"`
	Txt        string    `json:"txt"`
	Order      string    `json:"order"`
	UploadTime time.Time `orm:"auto_now_add;type(datetime)"  json:"uploadTime"`
	UpdateUid  string    `json:"updateUid"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:dianshi8@tcp(gz-cdb-c5qq3k9t.sql.tencentcdb.com:60981)/dianshi?charset=utf8")

	orm.RegisterModel(new(TbUser), new(TbAdUnit), new(TbResUnit), new(TbPic))

	//初始化ormer
	Db = orm.NewOrm()
	fmt.Println("dbbase open")
}
