package dto

type Budget struct {
	ID        int64  `json:"id"`
	Code      string `json:"code"`
	Debt      int    `json:"debt"`
	CreatedAt string `json:"created_at"`
}
