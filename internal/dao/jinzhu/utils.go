package jinzhu

import (
	"go-quick-template/internal/model"

	"gorm.io/gorm"
)

// 根据IDs获取用户列表
func getUsersByIDs(db *gorm.DB, ids []int64) ([]*model.User, error) {
	user := &model.User{}
	return user.List(db, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}
