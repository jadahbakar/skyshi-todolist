package activity

type Activity struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostReq struct {
	Title string `json:"title"`
	Email string `json:"email"`
}
