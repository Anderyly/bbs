/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package v1

import (
	"bbs/app/biz"
	"bbs/ay"
	"bbs/ay/lib"
	"bbs/dao/model"
	"github.com/gin-gonic/gin"
)

type Passport struct {
}

type register struct {
	Account        string `form:"account"`
	Password       string `form:"password"`
	RepeatPassword string `form:"repeat_password"`
}

func (con Passport) Register(c *gin.Context) {
	var err error
	output := lib.NewJson(c)
	var data register
	if err = c.ShouldBind(&data); err != nil {
		output.Fail(ay.Validator{}.Translate(err))
		return
	}
	b := biz.NewBiz(c)
	var user *model.User
	var token string
	if user, token, err = b.Passport.Register(data.Account, data.Password, data.RepeatPassword); err != nil {
		output.Fail(err.Error())
		return
	}

	output.Success(gin.H{
		"user":  user,
		"token": token,
	})
}

type login struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

func (con Passport) Login(c *gin.Context) {
	var err error
	output := lib.NewJson(c)
	var data login
	if err = c.ShouldBind(&data); err != nil {
		output.Fail(ay.Validator{}.Translate(err))
		return
	}
	b := biz.NewBiz(c)
	var user *model.User
	var token string
	if user, token, err = b.Passport.Login(data.Account, data.Password); err != nil {
		output.Fail(err.Error())
		return
	}

	output.Success(gin.H{
		"user":  user,
		"token": token,
	})
}
