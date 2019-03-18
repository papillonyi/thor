package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/internal/auth_service"
	"github.com/papillonyi/thor/pkg/app"
	"github.com/papillonyi/thor/pkg/e"
	"github.com/papillonyi/thor/pkg/util"
	"github.com/prometheus/common/log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")
	log.Info(username)
	log.Info(password)

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		//app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	session := sessions.Default(appG.C)
	session.Set("token", token)
	session.Save()
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
