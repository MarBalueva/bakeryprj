package models

// DictionaryItem общая структура элемента справочника
// @Description Универсальный элемент справочника (используется для всех типов справочников)
type DictionaryItem struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(256)"`
	IsDeleted bool   `json:"isdeleted" gorm:"default:false"`
}

type Position DictionaryItem
type AccessGroup DictionaryItem
type SubcategoryProduct DictionaryItem
type Status DictionaryItem
type PaymentType DictionaryItem
