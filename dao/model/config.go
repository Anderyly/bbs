package model

import "bbs/dao"

// Config 配置
type Config struct {
	dao.Model
	K        string `gorm:"not null;index" json:"k"`         // 键
	V        string `gorm:"not null;type:longtext" json:"v"` // 值
	Describe string `gorm:"not null" json:"describe"`        // 描述
}
