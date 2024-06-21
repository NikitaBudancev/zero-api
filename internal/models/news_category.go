package models

import _ "gopkg.in/reform.v1"

type NewsCategory struct {
	Id         int64 `reform:"id,pk"`
	NewsId     int64 `reform:"news_id"`
	CategoryId int64 `reform:"category_id"`
}

//go:generate reform
