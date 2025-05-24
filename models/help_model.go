package models

type AccessGroup struct {
	ID   int64  `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type ProductInBasket struct {
	ProductID int64 `gorm:"column:productid;primaryKey" json:"productid"`
	ClientID  int64 `gorm:"column:clientid;primaryKey" json:"clientid"`
	Count     int   `gorm:"column:count" json:"count"`
}

type ProductInOrder struct {
	ProductID int64   `gorm:"column:productid;not null;primaryKey" json:"productid"`
	OrderID   int64   `gorm:"column:orderid;not null;primaryKey" json:"orderid"`
	Count     int     `gorm:"column:count;not null" json:"count"`
	Cost      float64 `gorm:"column:cost;not null" json:"cost"`
}
