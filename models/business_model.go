package models

import "time"

type Product struct {
	//gorm.Model
	ID            int64   `gorm:"primaryKey;column:id" json:"id"`
	CategoryId    int64   `gorm:"column:categoryid" json:"categoryid"`
	Name          string  `gorm:"column:name" json:"name"`
	Description   string  `gorm:"column:description" json:"description"`
	Proteins      float64 `gorm:"column:proteins" json:"proteins"`
	Fats          float64 `gorm:"column:fats" json:"fats"`
	Carbohydrates float64 `gorm:"column:carbohydrates" json:"carbohydrates"`
	Calories      float64 `gorm:"column:calories" json:"calories"`
	UnWeight      float64 `gorm:"column:unweight" json:"unweight"`
	Weight        int     `gorm:"column:weight" json:"weight"`
	//Packaged      int     `gorm:"column:packaged" json:"packaged"`
	CountTypePack int     `gorm:"column:counttypepack" json:"counttypepack"`
	Cost          float64 `gorm:"column:cost" json:"cost"`
	//OkeiId        int     `gorm:"column:okeiid" json:"okeiId"`
	InStore   bool   `gorm:"column:instore" json:"instore"`
	PhotoLink string `gorm:"column:photolink" json:"photolink"`
	IsDeleted bool   `gorm:"column:is_deleted" json:"-"`
}

type PriceHistory struct {
	ProductID int64      `gorm:"column:productid;primaryKey" json:"productid"`
	StartDate time.Time  `gorm:"column:startdate;primaryKey" json:"startdate"`
	EndDate   *time.Time `gorm:"column:enddate" json:"enddate"`
	Cost      float64    `gorm:"column:cost" json:"cost"`
}

type Employee struct {
	ID         int64      `gorm:"primaryKey;column:id" json:"id"`
	Surname    string     `gorm:"column:surname" json:"surname"`
	Name       string     `gorm:"column:name" json:"name"`
	Patronymic *string    `gorm:"column:patronymic" json:"patronymic"`
	Phone      string     `gorm:"column:phonenumber" json:"phonenumber"`
	Email      string     `gorm:"column:email" json:"email"`
	StartDate  time.Time  `gorm:"column:startdate" json:"startdate"`
	EndDate    *time.Time `gorm:"column:enddate" json:"enddate"`
	PositionID int64      `gorm:"column:jobpositionid" json:"jobpositionid"`
	Number     string     `gorm:"column:number" json:"number"`
	PhotoLink  *string    `gorm:"column:photolink" json:"photolink"`
}

type Order struct {
	ID           int64      `gorm:"primaryKey;column:id" json:"id"`
	Name         string     `gorm:"column:name" json:"name"`
	RespEmpId    *int64     `gorm:"column:respempid" json:"respempid"`
	ClientId     int64      `gorm:"column:clientid" json:"clientid"`
	AddressId    *int64     `gorm:"column:addressid" json:"addressid"`
	StatusId     string     `gorm:"column:statusid" json:"statusid"`
	CreateDate   time.Time  `gorm:"column:createdate" json:"createdate"`
	SumOrder     float64    `gorm:"column:sumorder" json:"sumorder"`
	IsPay        bool       `gorm:"column:ispay" json:"ispay"`
	Comment      *string    `gorm:"column:comment" json:"comment"`
	EndDate      *time.Time `gorm:"column:enddate" json:"enddate"`
	DelStartDate *time.Time `gorm:"column:delstartdate" json:"delstartdate"`
	DelEndDate   *time.Time `gorm:"column:delenddate" json:"delenddate"`
	IsDeleted    bool       `gorm:"column:isDeleted" json:"isDeleted"`
}

type Client struct {
	ID          int64   `gorm:"primaryKey;column:id" json:"id"`
	Surname     string  `gorm:"column:surname" json:"surname"`
	Name        string  `gorm:"column:name" json:"name"`
	Patronymic  *string `gorm:"column:patronymic" json:"patronymic"`
	Email       string  `gorm:"column:email" json:"email"`
	PhoneNumber *string `gorm:"column:phonenumber" json:"phonenumber"`
	IsDeleted   bool    `gorm:"column:is_deleted;default:false" json:"isDeleted"`
}

type Payment struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	OrderID   int64     `gorm:"column:orderid;not null" json:"orderid"`
	Date      time.Time `gorm:"column:date;not null" json:"date"`
	Sum       float64   `gorm:"column:sum;not null" json:"sum"`
	PayTypeID int       `gorm:"column:paytypeid;not null" json:"paytypeid"`
}

type Document struct {
	ID         int64      `gorm:"primaryKey;column:id" json:"id"`
	Number     string     `gorm:"column:number;type:varchar(256)" json:"number"`
	Name       string     `gorm:"column:name;type:varchar(256)" json:"name"`
	CreateDate time.Time  `gorm:"column:createdate" json:"createDate"`
	StartDate  time.Time  `gorm:"column:startdate" json:"startDate"`
	EndDate    *time.Time `gorm:"column:enddate" json:"endDate"`
	FileLink   string     `gorm:"column:filelink;type:varchar(256)" json:"fileLink"`

	ProductID *int64 `gorm:"column:productid" json:"productId"`
	OrderID   *int64 `gorm:"column:orderid" json:"orderId"`
	LoadEmpID *int64 `gorm:"column:loadempid" json:"loadEmpId"`

	IsSubmit bool `gorm:"column:issubmit" json:"isSubmit"`
	Status   bool `gorm:"column:status" json:"status"`
}
