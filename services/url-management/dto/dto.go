package dto

type ShortenURLDTO struct {
	UserID      string
	RedirectTo  string `json:"redirect_to" binding:"required,url"`
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description" binding:"max=2000"`
}
