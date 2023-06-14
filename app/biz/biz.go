/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package biz

import (
	"github.com/gin-gonic/gin"
)

type Biz struct {
	User     *user
	Article  *article
	Passport *passport
}

type base struct {
	biz *Biz
	ctx *gin.Context
}

func NewBiz(ctx *gin.Context) *Biz {
	b := &Biz{}
	base := &base{biz: b, ctx: ctx}

	b.User = &user{base}
	b.Article = &article{base}
	b.Passport = &passport{base}

	return b
}
