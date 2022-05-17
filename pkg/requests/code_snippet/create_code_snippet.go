package code_snippet

import "time"

type CreateCodeSnippetRequest struct {
	LanguageId int       `json:"language_id" binding:"required"`
	Header     string    `json:"header" binding:"required,min=3,max=50"`
	Code       string    `json:"code" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" time_format:"2006-01-02" time_utc:"1"`
	ViewsLimit int       `json:"views_limit" binding:"omitempty,min=1"`
}
