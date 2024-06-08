package entity

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   string `json:"salary"`
}
