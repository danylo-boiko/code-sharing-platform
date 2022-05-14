package models

import "time"

type CodeSnippet struct {
	Id          int
	UserId      int
	LanguageId  int
	Header      string
	Code        string
	CreatedDate time.Time
	DeletedDate time.Time
	ViewsLimit  int
	Views       int
}
