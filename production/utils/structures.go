package utils

type Authors struct {
	ID int32 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type Users struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Prompt string `json:"prompt"`
}

// var user Users = Users{}

