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
	Uid        int       `orm:"column(uid);pk";json:"name"` // 设置主键
	RtxName    string    `json:"rtxName"`
	Password   string    `json:"password"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)";json:"createTime"`
}

type TbAdUnit struct {
	AdId          int    `orm:"column(ad_id);pk"` // 设置主键
	AdName        string `json:"adName"`
	AdSubTitle    string
	ResId         int
	Type          int
	Order         int
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`
	CreatorUid    int
	UpdateUid     int
	RelevantUsers string
	IsPublished   int
	Title         string
	BgType        int
	BgColor       string
	BgPicUrl      string
	LinkType      int
	ButtonName    string
	Link          string
	OnTime        time.Time
	OffTime       time.Time
	Remark        string
	IsUrgency     int
	UrgencyMsg    string
	IsUrgencyLink int
	UrgencyUrl    string
	ViewNum       int
}

type TbResUnit struct {
	ResId         int       `orm:"column(res_id);pk"` // 设置主键
	ResName       string    `json:"resName"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)"`
	CreatorUid    int
	UpdateUid     int
	RelevantUsers string
	IsPublished   int
	Count         int `orm:"-"`
	ResTitle      string
	LinkType      int
	ButtonName    string
	Link          string
	Url           string
	BgType        int
	BgColor       string
	BgPicUrl      string
	Msg           string
	Remark        string
	IsUrgency     int
	UrgencyMsg    string
	IsUrgencyLink int
	UrgencyUrl    string
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
