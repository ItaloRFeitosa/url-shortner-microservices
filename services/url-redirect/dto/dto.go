package dto

type AddShortURLDTO struct {
	RedirectTo string `json:"redirect_to"`
	ID         int    `json:"id"`
}

type RedirectDTO struct {
	Slug string `uri:"slug"`
}
