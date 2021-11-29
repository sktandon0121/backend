package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sktandon0121/backend/models"
	"github.com/sktandon0121/backend/repo"
	"github.com/sktandon0121/backend/services"
	"github.com/sktandon0121/backend/utils"
)

var userService = services.NewUserService(repo.NewUserRepo())
var errorRes = utils.NewCommonErrorResponse()

func Signup(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}

	signup := models.Signup{}
	if err := json.Unmarshal(body, &signup); err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}

	data, err := userService.SignUp(signup)
	if err != nil {
		WriteJsonData(rw, errorRes.InternalServerError(err, err.Error()), 500)
		return
	}

	WriteJsonData(rw, data, 201)
}

func Login(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}

	login := &models.LoginCredentials{}
	if err := json.Unmarshal(body, login); err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}

	data, err := userService.Login(login)
	if err != nil {
		WriteJsonData(rw, errorRes.InternalServerError(err, err.Error()), 500)
		return
	}
	WriteJsonData(rw, data, 200)
}

func Validate(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}

	token := &models.ValidateToken{}
	if err := json.Unmarshal(body, token); err != nil {
		WriteJsonData(rw, errorRes.BadRequest(err, err.Error()), 400)
		return
	}
	data, err := userService.Validate(token)
	validData := make(map[string]bool, 0)
	validData["valid"] = data
	if err != nil {
		WriteJsonData(rw, errorRes.InternalServerError(err, err.Error()), 500)
		return
	}
	WriteJsonData(rw, validData, 200)
}

func Buy(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Name   string
		UserId int
	}{Name: "Subodh", UserId: utils.GetUserFromContext(r.Context())}

	WriteJsonData(rw, data, 200)
}

func Sell(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]string, 0)
	data["name"] = "Subodh Tandon"
	WriteJsonData(rw, data, 200)
}
