package user

type User struct {
	TelegramID       int64
	TelegramUsername string
	Role             UserRole
}

type UserRole string

var (
	RoleAdmin   UserRole = "admin"
	RoleUser    UserRole = "user"
	RoleManager UserRole = "manager"
)
