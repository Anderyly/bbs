/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package services

type Pager struct {
	Limit  int
	Offset int
}

func NewPage(page, limit int) *Pager {
	if limit > 20 {
		limit = 20
	}
	offset := (page - 1) * limit
	return &Pager{
		Limit:  limit,
		Offset: offset,
	}
}

func (con Pager) GetOffset() int {
	return con.Offset
}

func (con Pager) GetLimit() int {
	return con.Limit
}
