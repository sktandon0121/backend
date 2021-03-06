package repo

import (
	"log"

	"github.com/sktandon0121/backend/models"
)

type UserRepo interface {
	SaveUser(user *models.User) (*models.User, error)
	FindUserByUserName(username string, preload bool) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	FindUserByUserId(userId int) (*models.User, error)
	UpdateWallet(userId int, wallet float64) error
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
		if res := db.Preload("Bitcoin").Preload("Wallet").Where("user_name", username).Find(&user); res.Error != nil {
			return nil, res.Error
		}
	} else {
		if res := db.Find(&user); res.Error != nil {
			return nil, res.Error
		}
	}

	return user, nil
}

func (u *userRepo) UpdateUser(user *models.User) (*models.User, error) {
	db := GormDB()
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) FindUserByUserId(userId int) (*models.User, error) {
	db := GormDB()
	var user *models.User
	result := db.Preload("Bitcoin").Preload("Wallet").Where("id", userId).Find(&user)

	if result.Error != nil {
		log.Print("Error in getting user by id ", result.Error)
		return user, result.Error
	}
	return user, nil
}

func (u *userRepo) UpdateWallet(userId int, wallet float64) error {
	db := GormDB()
	if err := db.Table("wallets").Where("user_id = ? ", userId).Update("value", wallet).Error; err != nil {
		log.Printf("Error while updating the wallet : %#v", err)
		return err
	}
	return nil
}
