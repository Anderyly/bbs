package model

import "bbs/dao"

// UserLoginLog 登入日志
type UserLoginLog struct {
	dao.Model
	UserId     int    `gorm:"not null;index"`            // 账号
	RegisterIp string `gorm:"not null;type:varchar(15)"` // 注册ip
	LoginIp    string `gorm:"not null;type:varchar(15)"` // 登入ip
}
