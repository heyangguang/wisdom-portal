package models

type SwaggerUser struct {
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
