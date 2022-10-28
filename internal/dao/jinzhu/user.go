package jinzhu

import (
	"strings"

	"go-quick-template/internal/core"
	"go-quick-template/internal/model"

	"gorm.io/gorm"
)

var (
	_ core.UserManageService = (*userManageServant)(nil)
)

type userManageServant struct {
	db *gorm.DB
}

func newUserManageService(db *gorm.DB) core.UserManageService {
	return &userManageServant{
		db: db,
	}
}

func (s *userManageServant) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{
		Model: &model.Model{
			ID: id,
		},
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{
		Email: email,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{
		Phone: phone,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUsersByIDs(ids []int64) ([]*model.User, error) {
	user := &model.User{}
	return user.List(s.db, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *userManageServant) GetUsersByKeyword(keyword string) ([]*model.User, error) {
	user := &model.User{}
	keyword = strings.Trim(keyword, " ") + "%"
	if keyword == "%" {
		return user.List(s.db, &model.ConditionsT{
			"ORDER":       "id ASC",
			"is_set != ?": 0,
		}, 0, 6)
	} else {
		return user.List(s.db, &model.ConditionsT{
			"username LIKE ?": keyword,
			"is_set != ?":     0,
		}, 0, 6)
	}
}

func (s *userManageServant) CreateUser(user *model.User) (*model.User, error) {
	return user.Create(s.db)
}

func (s *userManageServant) UpdateUser(user *model.User) error {
	return user.Update(s.db)
}

func (s *userManageServant) IsFriend(userId int64, friendId int64) bool {
	// just true now
	return true
}
