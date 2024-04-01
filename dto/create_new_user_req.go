package dto

type CreateNewUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
