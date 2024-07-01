package dbo

import (
	"strings"

	"gorm.io/gorm"
)

type UserInfo struct {
	UserID   string
	UserName string
	// UserPassword string
	// TipId        string
	LableId string
	Email   string
}

// 查询所有的数据
func GetAllUserInfo(db *gorm.DB) ([]map[string]interface{}, error) {
	var users []UserInfo
	result := make([]map[string]interface{}, 0)

	if err := db.Find(&users).Error; err != nil {
		return result, err
	}
	for _, user := range users {
		jsonObject := make(map[string]interface{})
		jsonObject["user_id"] = strings.TrimSpace(user.UserID)
		jsonObject["user_name"] = strings.TrimSpace(user.UserName)
		jsonObject["lable_id"] = strings.TrimSpace(user.LableId)
		jsonObject["email"] = strings.TrimSpace(user.Email) // 修改为正确的字段名
		result = append(result, jsonObject)
	}
	return result, nil
}
