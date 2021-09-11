package users

type UserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type User struct {
	Id        int
	Email     string
	Password  string
	CreatedAt int
	UpdatedAt int
}
