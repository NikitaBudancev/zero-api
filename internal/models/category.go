package models

import _ "gopkg.in/reform.v1"

type Category struct {
	Id   int64  `reform:"id,pk"`
	Name string `reform:"name"`
}

//go:generate reform
