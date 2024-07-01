package dbo

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type DataBat struct {
	Name     string
	ImageUrl string
	VideoUrl string
	ImageId  string
	VideoId  string
}

type DataBats struct {
	Id       uint
	Name     string
	ImageUrl string
	VideoUrl string
	ImageId  string
	VideoId  string
}

func (DataBats) TableName() string {
	return "data_bat"
}

// 中间件 验证 必填字段
func middlewareFields(data map[string]string, method string) error {
	var requiredFields []string

	if method == "add" {
		requiredFields = []string{"name"}
	}

	if method == "update" {
		requiredFields = []string{"id", "name"}
	}

	for _, field := range requiredFields {
		if _, ok := data[field]; !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}
	return nil
}

// 单条新增
func AddOneDataInfo(db *gorm.DB, data map[string]string) error {
	err := middlewareFields(data, "add")
	if err != nil {
		return err
	}

	info := DataBat{
		Name:     data["name"],
		ImageUrl: data["image_url"],
		VideoUrl: data["video_url"],
		ImageId:  data["image_id"],
		VideoId:  data["video_id"],
	}

	if err := db.Create(&info).Error; err != nil {
		return err
	}
	return nil
}

// 单条修改
func UpdateOneDataInfo(db *gorm.DB, data map[string]string) error {
	// db.AutoMigrate(&DataBats{})
	err := middlewareFields(data, "update")
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	info := &DataBats{
		Id:       uint(id),
		Name:     data["name"],
		ImageUrl: data["image_url"],
		VideoUrl: data["video_url"],
		ImageId:  data["image_id"],
		VideoId:  data["video_id"],
	}

	if err := db.Save(&info).Error; err != nil {
		return err
	}
	return nil
}
