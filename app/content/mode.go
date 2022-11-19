package content

import (
	"cmsGo/app/common"
	"log"

	"gorm.io/gorm"
)

type Content struct {
	// 结构体嵌套 这里提前定义了：id，createTime，endTime
	gorm.Model
	Work
}

// content 表的业务逻辑字段
type Work struct {
	Subject string `gorm:"type:varchar(255);index" json:"subject"`
	Summary string `gorm:"type:varchar(255)" json:"summary"`
	Body    string `gorm:"type:text" json:"body"`
	UserID  uint   `gorm:"index" json:"user_id"`
	Status  string `gorm:"type:enum('new','draft','publish')" json:"status"`
}

// init
func init() {
	// 自动创建管理相关的表
	if err := common.DB.AutoMigrate(&Content{}); err != nil {
		log.Println(err)
	}
}
