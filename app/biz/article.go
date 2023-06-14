/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package biz

import (
	"bbs/app/services"
	"bbs/ay"
	"bbs/dao/model"
	"bbs/dao/query"
)

type article struct {
	*base
}

func (con article) List(plateId int, pager *services.Pager) (res []*model.Article, count int64, err error) {
	qe := query.Use(ay.Db).Article
	do := qe.WithContext(con.ctx).Where(qe.Plate.Eq(plateId))

	if count, err = do.Count(); err != nil {
		return
	}

	if res, err = do.
		Offset(pager.GetOffset()).
		Limit(pager.GetLimit()).
		Find(); err != nil {
		return
	}

	return
}
