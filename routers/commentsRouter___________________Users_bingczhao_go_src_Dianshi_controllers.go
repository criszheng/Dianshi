package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Dianshi/controllers:AdUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:AdUnitController"],
        beego.ControllerComments{
            Method: "ClickAd",
            Router: `/clickAd`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:AdUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:AdUnitController"],
        beego.ControllerComments{
            Method: "UpdateOrder",
            Router: `/updateOrder`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:PicController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:PicController"],
        beego.ControllerComments{
            Method: "SavePicture",
            Router: `/savePicture`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"],
        beego.ControllerComments{
            Method: "AddResUnit",
            Router: `/addResUnit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"],
        beego.ControllerComments{
            Method: "DeleteResUnit",
            Router: `/deleteResUnit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"],
        beego.ControllerComments{
            Method: "GetAllResUnit",
            Router: `/getAllResUnit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"],
        beego.ControllerComments{
            Method: "GetResUnit",
            Router: `/getResUnit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"] = append(beego.GlobalControllerRouter["Dianshi/controllers:ResUnitController"],
        beego.ControllerComments{
            Method: "UpdateResUnit",
            Router: `/updateResUnit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
