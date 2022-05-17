package code_snippet

import "time"

type UpdateCodeSnippetRequest struct {
	LanguageId int
	Header     string
	Code       string
	ExpiryDate time.Time
	ViewsLimit int
	Views      int
}
