package server

import (
	"bbs/ay"
	"bbs/ay/lib"
	"bbs/dao/model"
	"bbs/dao/query"
	"errors"
	"github.com/gin-gonic/gin"
)

type Common struct {
}

var (
	Token  string
	Expiry int64
)

func init() {
	Expiry = 3600 * 2
}

// MustLogon 判断登录
func (con Common) MustLogon(ctx *gin.Context) (user *model.User, err error) {
	qe := query.Use(ay.Db).User
	str := lib.NewStr()
	if user, err = ay.IgnoreNotFoundReturn(qe.WithContext(ctx).Where(qe.Account.Eq(str.AuthCode(Token, "DECODE", ay.Yaml.GetString("key"), Expiry))).Take()); err != nil {
		return
	}

	if user == nil {
		err = errors.New("must be login")
		return
	}

	return
}
