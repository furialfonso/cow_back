package response

type UsersByTeamResponse struct {
	BudgetName string         `json:"name_budget"`
	Users      []UserResponse `json:"users"`
}

type TeamsByUserResponse struct {
	Teams []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	Code string `json:"code"`
	Debt int    `json:"debt"`
}
