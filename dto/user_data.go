package dto

type UserData struct {
	ID       string `json:"id"`
	Email    string `json:"name"`
	Password string `json:"-"`
}
