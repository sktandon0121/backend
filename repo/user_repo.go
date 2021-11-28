package repo

import "github.com/sktandon0121/backend/models"

type UserRepo interface {
	SaveUser(user *models.User) (*models.User, error)
	FindUserByUserName(username string, preload bool) (*models.User, error)
}
type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) SaveUser(user *models.User) (*models.User, error) {
	db := GormDB()
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) FindUserByUserName(username string, preload bool) (*models.User, error) {
	db := GormDB()
	user := &models.User{}
	if preload {
		if res := db.Preload("Bitcoin").Preload("Wallet").Find(&user); res.Error != nil {
			return nil, res.Error
		}
	} else {
		if res := db.Find(&user); res.Error != nil {
			return nil, res.Error
		}
	}

	return user, nil
}