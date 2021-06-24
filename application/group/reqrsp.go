package group

type GetGroupRequest struct {
	Id string `json:"id"`
}

type GetGroupResponse struct {
	Description string `json:"description"`
}
