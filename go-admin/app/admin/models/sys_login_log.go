package models

import (
	"encoding/json"
	"errors"
	"go-admin/common/models"
	log "go-admin/core/logger"
	"go-admin/core/sdk"
	"go-admin/core/storage"
	"time"
)

type SysLoginLog struct {
	models.Model
	models.ControlBy
	Username      string    `json:"username" gorm:"size:128;comment:用户名"`
	Status        string    `json:"status" gorm:"size:4;comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"size:255;comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"size:255;comment:归属地"`
	Browser       string    `json:"browser" gorm:"size:255;comment:浏览器"`
	Os            string    `json:"os" gorm:"size:255;comment:系统"`
	Platform      string    `json:"platform" gorm:"size:255;comment:固件"`
	Remark        string    `json:"remark" gorm:"size:255;comment:备注"`
	Msg           string    `json:"msg" gorm:"size:255;comment:信息"`
	LoginTime     time.Time `json:"loginTime" gorm:"comment:登录时间"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

func (e *SysLoginLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysLoginLog) GetId() interface{} {
	return e.Id
}

// SaveLoginLog 从队列中获取登录日志
func SaveLoginLog(message storage.Messager) (err error) {
	// 准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		return err
	}

	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		return err
	}

	var sysLoginLog SysLoginLog
	err = json.Unmarshal(rb, &sysLoginLog)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}

	err = db.Create(&sysLoginLog).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		return err
	}
	
	return nil
}