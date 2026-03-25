package users

type User struct {
	UID      string
	Name     string
	Age      int
	Email    string
	Password string
	Role     Role
	Balance  float64
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=18"` // gte === >=
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
