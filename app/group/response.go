package group

type GetGroupResponse struct {
	ID            string `json:"id"`
	Privacy       string `json:"privacy"`
	OwnerID       string `json:"ownerId"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Avatar        string `json:"avatar"`
	Cover         string `json:"cover"`
	Description   string `json:"description"`
	Terms         string `json:"terms"`
	MemberCount   int    `json:"memberCount"`
	Deleted       bool   `json:"deleted"`
	JoinByDefault bool   `json:"joinByDefault"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}
