package users

type Role int

const (
	AdminRole = iota
	OwnerRole
	UserRole
)

var roles = []string{"admin", "owner", "user"}

func (r Role) String() string {
	return roles[r]
}

func ParseRole(role string) Role {
	switch role {
	case "admin":
		return AdminRole
	case "owner":
		return OwnerRole
	default:
		return UserRole
	}
}
