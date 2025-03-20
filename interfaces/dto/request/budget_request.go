package request

type BudgetRequest struct {
	Code string `json:"name_budget"`
	Debt int    `json:"debt"`
}
