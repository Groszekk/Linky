package models

type Short struct {
	Link  string `json:"link" validate:"min=4,max=8000"`
	Short string
}
