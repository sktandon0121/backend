package controllers

import (
	"net/http"
	"strings"

	"github.com/sktandon0121/backend/utils"
)

type midSvc struct{}

func AuthMiddleware() *midSvc {
	return &midSvc{}
}

func (m *midSvc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := extractToken(r)
	isValid, err := utils.VerifyToken(token)
	if err != nil {
		WriteJsonData(rw, errorRes.AccessDenied(err, "Access denied"), 401)
		return
	}

	// Get data from the token
	data, err1 := utils.GetTokenMetadata(token)
	if err != nil {
		WriteJsonData(rw, errorRes.AccessDenied(err1, "Access denied"), 401)
		return
	}
	// set user id to request context
	ctx := utils.AddDataToContext(r.Context(), data, token)
	r = r.WithContext(ctx)
	if isValid {
		next(rw, r)
	} else {
		WriteJsonData(rw, errorRes.AccessDenied(err, "Access denied"), 401)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
