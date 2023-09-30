package todos

type Todo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	ActivityId int    `json:"activity_group_id"`
	IsActive   bool   `json:"is_active"`
	Priority   string `json:"priority"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type PostReq struct {
	Title      string `json:"title"`
	ActivityId int    `json:"activity_group_id"`
	IsActive   bool   `json:"is_active"`
}

type PatchReq struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
}
