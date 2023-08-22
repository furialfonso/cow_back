package response

type TeamUsersResponse struct {
	GroupName string         `json:"name_group"`
	Users     []UserResponse `json:"users"`
}
