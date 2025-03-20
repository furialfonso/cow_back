package dto

type Team struct {
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	BudgetID  int64  `json:"budget_id"`
}
