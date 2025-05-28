package models

// type AccessGroup struct {
// 	ID   int64  `gorm:"primaryKey;column:id" json:"id"`
// 	Name string `gorm:"column:name" json:"name"`
// }

type ProductInBasket struct {
	ProductID int64 `gorm:"column:productid;primaryKey" json:"productid" binding:"required,gt=0"`
	ClientID  int64 `gorm:"column:clientid;primaryKey" json:"clientid" binding:"required,gt=0"`
	Count     int   `gorm:"column:count" json:"count" binding:"required,gt=0"`
}

type ProductInOrder struct {
	ProductID int64   `gorm:"column:productid;not null;primaryKey" json:"productid"`
	OrderID   int64   `gorm:"column:orderid;not null;primaryKey" json:"orderid"`
	Count     int     `gorm:"column:count;not null" json:"count"`
	Cost      float64 `gorm:"column:cost;not null" json:"cost"`
}

type UserAccess struct {
	UserID  int64 `gorm:"column:userid;primaryKey" json:"userid"`
	GroupID int64 `gorm:"column:groupid;primaryKey" json:"groupid"`
}

type StatusInput struct {
	StatusId int64 `json:"statusid" binding:"required,gt=0"`
}

type StatusResponse struct {
	StatusId int64 `json:"statusid"`
}
