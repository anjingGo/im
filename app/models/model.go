package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
}

// CommonTimestampsField CreatedAt , UpdatedAt 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index" json:"updated_at,omitempty"`
}

// DeletedAtTimestampsField DeletedAt 时间戳
type DeletedAtTimestampsField struct {
	DeletedAt time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}
