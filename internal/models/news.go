package models

import _ "gopkg.in/reform.v1"

type News struct {
	Id      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}

//go:generate reform
