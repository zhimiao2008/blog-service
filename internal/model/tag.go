package model

import (
	"github.com/go-grogramming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

type Tag struct {
	*Model
	// 标签名称
	Name string `json:"name"`
	// 标签状态， 0 为禁用， 1 为启用
	State uint8 `json:"state"`
}

func (m Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ? ", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//return db.Model(&t).Where("id = ? and is_del = ?", t.ID, 0).Updates(values).Error
	if err := db.Model(t).Updates(values).Where("id = ? and is_del = ?", t.ID, 0).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? ", t.ID).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}