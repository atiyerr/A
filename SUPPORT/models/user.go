// user数据表
package models

type User struct {
	Id       int    `gorm:"primary_key"` // 主键标签
	Username string `gorm:"not null"`    // 不能为空标签
	Identity string `gorm:"not null"`    // 不能为空标签
	Motto    string
	JoinDate string `gorm:"type:date"` // 日期类型标签
	Photo    string
}

func (User) TableName() string {
	return "user"
}
