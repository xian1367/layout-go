// Package model 模型通用属性和方法
package model

import (
	"gorm.io/gorm"
)

// BaseModel 模型基类
type BaseModel struct {
	IDField
	CommonTimestampsField
}

type IDField struct {
	ID int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // ID

}

type CommonTimestampsField struct {
	CreatedAt Carbon         `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt Carbon         `gorm:"column:updated_at" json:"updated_at"` // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
}

type Avg struct {
	Count *int32   `gorm:"->" json:"count"`
	Sum   *float64 `gorm:"->" json:"sum"`
	Max   *float64 `gorm:"->" json:"max"`
	Min   *float64 `gorm:"->" json:"min"`
}
