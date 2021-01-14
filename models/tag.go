package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model

	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	State     int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetCount(maps interface{}) (count int) {
	db.Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createBy string) bool {
	// 就下面这样，万一插入错误如何处理？？？
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})

	return true
}
