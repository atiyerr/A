// load数据表
package models

type Load struct {
	Password string `gorm:"not null"`                    // 不能为空标签
	Userid   string `gorm:"unique;not null;primary_key"` // 唯一且不能为空标签
}

func (Load) TableName() string {
	return "load"
}
