package userDTO

type CreateUserRequestDTO struct {
	Params map[string]interface{}
}

type GetUserRequestDTO struct {
	UserID int
}

type UpdateUserRequestDTO struct {
	UserID int
	Params map[string]interface{}
}

type DeleteUserRequestDTO struct {
	UserID int
}
