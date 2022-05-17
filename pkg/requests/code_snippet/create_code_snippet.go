package code_snippet

import "time"

type CreateCodeSnippetRequest struct {
	LanguageId int
	Header     string
	Code       string
	ExpiryDate time.Time
	ViewsLimit int
	Views      int
}
