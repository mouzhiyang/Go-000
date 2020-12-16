package service

import (
	"../model"
)

func QueryDevice() (*[]model.Device, error) {
	devices := []model.Device{}
	err := db.Table("device").
		Select("*").
		Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return &devices, nil
}
