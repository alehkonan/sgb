package repository

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	Ru string `db:"ru_word" json:"russianWord"`
	Ka string `db:"ka_word" json:"georgianWord"`
}
