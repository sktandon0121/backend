package services

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/sktandon0121/backend/models"
	"github.com/sktandon0121/backend/repo"
	"github.com/sktandon0121/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(data models.Signup) (*models.User, error)
	Login(loginCred *models.LoginCredentials) (*models.LoginResponse, error)
	Validate(token *models.ValidateToken) (bool, error)
	Buy(ctx context.Context) (*models.User, error)
}

type userSvc struct {
	userRepo repo.UserRepo
}

func NewUserService(uRepo repo.UserRepo) UserService {
	return &userSvc{userRepo: uRepo}
}

func (uSvc *userSvc) SignUp(data models.Signup) (*models.User, error) {

	if data.Password != data.ConfirmPassword {
		return nil, errors.New("confirm password and password are not same")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	if err != nil {
		return nil, err
	}

	wallet := &models.Wallet{Value: 500000}
	bitcoin := &models.Bitcoin{}
	user := &models.User{
		UserName: data.UserName,
		Password: string(pass),
		Wallet:   wallet,
		Bitcoin:  bitcoin,
	}

	// save to database
	savedUser, err := uSvc.userRepo.SaveUser(user)

	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

func (uSvc *userSvc) Login(loginCred *models.LoginCredentials) (*models.LoginResponse, error) {

	user, err := uSvc.userRepo.FindUserByUserName(loginCred.UserName, true)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid login credentials")
	}

	pass := strings.Trim(loginCred.Password, "\n")

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return nil, errors.New("invalid login credentials")
	}
	token, err := utils.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}
	loginRes := &models.LoginResponse{User: user, Token: token}

	return loginRes, nil
}

func (uSvc *userSvc) Validate(token *models.ValidateToken) (bool, error) {
	isValid, err := utils.VerifyToken(token.Token)
	if err != nil {
		return false, err
	}
	return isValid, nil
}

func (uSvc *userSvc) Buy(ctx context.Context) (*models.User, error) {
	// data := struct{
	// 	CurrentPrice float64
	// 	Amount float64
	// 	}{CurrentPrice: 300, Amount: 100}

	// userDetails := uSvc.userRepo.FindUserByUserId(userId)
	// balanceUsed := data.Amount * data.CurrentPrice
	// remainingBalance :=
	userId := utils.GetUserFromContext(ctx)

	walletValue := 490000
	err := uSvc.userRepo.UpdateWallet(userId, float64(walletValue))
	if err != nil {
		log.Println("Error in updating the wallet ", err)
	}
	return uSvc.userRepo.FindUserByUserId(userId)
}

func (uSvc *userSvc) Sell() {

}
