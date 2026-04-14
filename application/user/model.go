	package user

	type User struct {
		TelegramID int64
		Role       UserRole
	}

	type UserRole string
