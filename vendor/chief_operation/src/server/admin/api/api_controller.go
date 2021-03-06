package api

import (
	"chief_operation/src/server/admin/dao"
	"chief_operation/src/server/admin/entity"
	"chief_operation/src/server/admin/handler"
	"chief_operation/src/util/common"
	log "github.com/astaxie/beego/logs"
	"time"
)

type ApiController struct {}



func NewApiController() *ApiController {
	d := new(ApiController)
	return d
}

func (this *ApiController) PushMsg(param *handler.Param) error {
	log.Info("推送消息：" + common.Obj2JsonStr(param) )
	//消息推送到钉钉
	go handler.Add(param)

	var alarmMsg = new(entity.AlarmMsg)
	alarmMsg.AppName = param.AppName
	alarmMsg.CreateTime = time.Now()
	alarmMsg.StartTime = time.Now()
	alarmMsg.ErrMsg = param.Msg
	alarmMsg.Org = param.Org
	alarmMsg.Title = param.Error
	alarmMsg.ModuleName = param.ModuleName
	//写数据到数据库
	err := dao.MysqlCon.Save(alarmMsg).Error
	if err != nil {
		log.Error("写入消息失败"+err.Error())
		return err
	}
	return nil
}
