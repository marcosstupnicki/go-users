package users

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type User struct {
	ID        int    `gorm:"column:id;primaryKey"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}
