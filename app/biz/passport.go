/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package biz

import (
	"bbs/app/controllers/server"
	"bbs/ay"
	"bbs/ay/lib"
	"bbs/dao"
	"bbs/dao/model"
	"bbs/dao/query"
	"errors"
	"strconv"
)

type passport struct {
	*base
}

func (con passport) Register(account, password, repeatPassword string) (user *model.User, token string, err error) {
	if password != repeatPassword {
		err = errors.New("两次密码输入不一致")
		return
	}

	if user, err = con.biz.User.GetUserByAccount(account); err != nil {
		return
	}
	if user != nil {
		err = errors.New("用户已存在")
		return
	}

	user = &model.User{
		Account:  account,
		Password: lib.NewStr().Md5(password + ay.Yaml.GetString("key")),
		RoleId:   1,
		Status:   dao.StatusNormal,
	}
	qe := query.Use(ay.Db).User
	if err = qe.WithContext(con.ctx).Create(user); err != nil {
		return
	}

	if err = con.addLoginLog(user.Id, con.ctx.ClientIP(), false); err != nil {
		return
	}

	token = lib.NewStr().AuthCode(strconv.Itoa(user.Id), "ENCODE", ay.Yaml.GetString("key"), server.Expiry)
	return
}

func (con passport) Login(account, password string) (user *model.User, token string, err error) {

	if user, err = con.biz.User.GetUserByAccount(account); err != nil {
		return
	}
	if user == nil {
		err = errors.New("用户不存在")
		return
	}

	if user.Password != lib.NewStr().Md5(password+ay.Yaml.GetString("key")) {
		err = errors.New("密码错误")
		return
	}

	if err = con.addLoginLog(user.Id, con.ctx.ClientIP(), true); err != nil {
		return
	}

	token = lib.NewStr().AuthCode(strconv.Itoa(user.Id), "ENCODE", ay.Yaml.GetString("key"), server.Expiry)
	return
}

func (con passport) addLoginLog(userId int, ip string, isLogin bool) (err error) {
	qe := query.Use(ay.Db).UserLoginLog
	data := &model.UserLoginLog{
		UserId: userId,
	}
	if isLogin {
		data.LoginIp = ip
	} else {
		data.RegisterIp = ip
	}
	if err = qe.WithContext(con.ctx).Create(data); err != nil {
		return
	}
	return
}
