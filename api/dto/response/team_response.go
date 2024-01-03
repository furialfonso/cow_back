package response

type UsersByTeamResponse struct {
	GroupName string         `json:"name_group"`
	Users     []UserResponse `json:"users"`
}

type TeamsByUserResponse struct {
	Teams []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	Code string `json:"code"`
	Debt int    `json:"debt"`
}
