package models

// 用户表模型层

import (
	
)

type User struct {
	ID        int64 `gorm:"column:user_id;primary_key;auto_increment"`
	Name      string `gorm:"column:user_name;type:varchar(100);not null;default:''"`
}