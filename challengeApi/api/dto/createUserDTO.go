package dto

type CreateUserInputDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type CreatedUserOutputDTO struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
