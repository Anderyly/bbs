package model

import "bbs/dao"

// User 用户
type User struct {
	dao.Model
	Account  string     `gorm:"not null;index;type:varchar(20)" json:"account"` // 账号
	Email    string     `gorm:"not null;index;type:varchar(50)" json:"email"`   // 邮箱
	Mobile   string     `gorm:"not null;type:varchar(11)" json:"mobile"`        // 手机号
	RoleId   int        `gorm:"noy null" json:"-"`                              // 角色id
	Password string     `gorm:"not null;type:varchar(64)" json:"-"`             // 密码
	Status   dao.Status `gorm:"not null" json:"status"`                         // 状态
}
