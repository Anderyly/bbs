package dao

type Model struct {
	Id        int   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt int64 `gorm:"not null" json:"created_at"`
	UpdatedAt int64 `gorm:"not null" json:"-"`
	DeletedAt int64 `gorm:"not null" json:"-"`
}

type Status int

const (
	// 正常
	StatusNormal Status = iota + 1
	// 关闭
	StatusClose
)
