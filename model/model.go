package model

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	Wind  int `json:"wind"`
	Water int `json:"water"`
}
