/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package v1

import (
	"bbs/app/biz"
	"bbs/app/controllers/server"
	"bbs/app/services"
	"bbs/ay"
	"bbs/ay/lib"
	"bbs/dao/model"
	"github.com/gin-gonic/gin"
)

type Article struct {
	server.Common
}

type articleList struct {
	Page    int `form:"page"`
	Limit   int `form:"limit"`
	PlateId int `form:"plate_id"`
}

func (con Article) List(c *gin.Context) {
	var err error
	output := lib.NewJson(c)

	var data articleList
	if err = c.ShouldBind(&data); err != nil {
		output.Fail(ay.Validator{}.Translate(err))
		return
	}

	b := biz.NewBiz(c)

	pager := services.NewPage(data.Page, data.Limit)

	var res []*model.Article
	var count int64
	if res, count, err = b.Article.List(data.PlateId, pager); err != nil {
		output.Fail(err.Error())
		return
	}

	output.Success(gin.H{
		"list":  res,
		"count": count,
	})

}
