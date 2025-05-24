package models

import "time"

type Appuser struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Login      string    `gorm:"column:login;size:32" json:"login"`
	Password   string    `gorm:"column:password;size:2000" json:"password"`
	EmpId      *int64    `gorm:"column:empId" json:"empId"`
	ClientId   *int64    `gorm:"column:clientId" json:"clientId"`
	CreateDate time.Time `gorm:"column:createDate" json:"-"`
	IsActive   bool      `gorm:"column:isActive" json:"isActive"`
	IsDeleted  bool      `gorm:"column:is_deleted;default:false" json:"is_deleted"`
}
