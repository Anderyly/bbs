/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package biz

import (
	"bbs/ay"
	"bbs/dao/model"
	"bbs/dao/query"
)

type user struct {
	*base
}

func (con user) GetUserByAccount(account string) (user *model.User, err error) {
	qe := query.Use(ay.Db).User
	if user, err = qe.WithContext(con.ctx).Where(qe.Account.Eq(account)).Take(); err != nil {
		return
	}

	return
}
