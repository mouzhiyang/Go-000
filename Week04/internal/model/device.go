package model

type Device struct {
	ID int64 `json:"id" gorm:"column:id;primary_key"`
}
