package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetCount(maps interface{}) (count int) {
	db.Where(maps).Count(&count)

	return
}