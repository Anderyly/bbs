/*
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2023
 */

package model

import "bbs/dao"

// Article 文章
type Article struct {
	dao.Model
	Plate      int    `gorm:"not null;index" json:"-"`                 // 板块Id
	UserId     int    `gorm:"not null;index" json:"user_id"`           // 用户Id
	Title      string `gorm:"not null;type:varchar(255)" json:"title"` // 标题
	Content    string `gorm:"not null;type:longtext" json:"content"`   // 内容
	LikesCount int64  `gorm:"-" json:"likes_count"`
	Avatar     string `gorm:"-" json:"avatar"`
	Nickname   string `gorm:"-" json:"nickname"`
}

// ArticleLike 点赞
type ArticleLike struct {
	dao.Model
	UserId    int    `gorm:"not null;uniqueIndex:idx_likes_id" json:"user_id"`    // 用户Id
	ArticleId int    `gorm:"not null;uniqueIndex:idx_likes_id" json:"article_id"` // 文章id
	Avatar    string `gorm:"-" json:"avatar"`
	Nickname  string `gorm:"-" json:"nickname"`
}

// ArticleComment 评论
type ArticleComment struct {
	dao.Model
	UserId        int    `gorm:"not null;index" json:"user_id"`    // 用户Id
	ArticleId     int    `gorm:"not null;index" json:"article_id"` // 文章id
	ReplyId       int    `gorm:"not null;index" json:"reply_id"`   // 回复id
	Content       string `gorm:"not null" json:"content"`          // 内容
	Avatar        string `gorm:"-" json:"avatar"`
	Nickname      string `gorm:"-" json:"nickname"`
	ReplyNickname string `gorm:"-" json:"reply_nickname"`
	ReplyAvatar   string `gorm:"-" json:"reply_avatar"`
}
