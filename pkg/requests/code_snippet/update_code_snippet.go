package code_snippet

import "time"

type UpdateCodeSnippetRequest struct {
	LanguageId int       `json:"language_id"`
	Header     string    `json:"header" binding:"omitempty,min=3,max=50"`
	Code       string    `json:"code"`
	ExpiryDate time.Time `json:"expiry_date" time_format:"2006-01-02"`
	ViewsLimit int       `json:"views_limit" binding:"omitempty,min=1"`
}
