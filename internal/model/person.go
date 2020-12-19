package model

type Person struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type (
	GetPersonRequest struct {
		ID string `json:"id"`
	}
	GetPersonResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	GetPersonsRequest struct {
		Page    int32
		Limit   int32
		OrderBy []string
	}
)
