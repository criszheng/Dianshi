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
	Uid        int `orm:"column(uid);pk"` // 设置主键
	RtxName    string
	Password   string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

type TbAdUnit struct {
	AdID          int `orm:"column(ad_id);pk"` // 设置主键
	AdName        string
	AdSubname     string
	ResId         int
	Type          int
	Order         int
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`
	CreatorUid    int
	UpdateUid     int
	RelevantUsers string
	IsPublished   int
	ViewNum       int
}

type TbResUnit struct {
	ResId         int `orm:"column(res_id);pk"` // 设置主键
	ResName       string
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`
	CreatorUid    int
	UpdateUid     int
	RelevantUsers string
	IsPublished   int
	Count         int `orm:"-"`
}

type TbPic struct {
	PicId      int `orm:"column(pic_id);pk"` // 设置主键
	AdId       int
	Url        string
	Link       string
	Txt        string
	Order      string
	UploadTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateUid  int
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:dianshi8@tcp(gz-cdb-c5qq3k9t.sql.tencentcdb.com:60981)/dianshi?charset=utf8")

	orm.RegisterModel(new(TbUser), new(TbAdUnit), new(TbResUnit), new(TbPic))

	//初始化ormer
	Db = orm.NewOrm()
	fmt.Println("dbbase open")
}
