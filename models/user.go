package models

import "time"

// user struct to save user on database
type User struct {
	ID        int64     `gorm:primary_key;auto_increment;json:id`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at",omitempty`
	UpdatedAt time.Time `json:"updated_at",omitempty`
}

// return table name of user model
func (user *User) TableName() string {
	return "user"
}

// request binding for user login
type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// request binding for user register
type UserRegister struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

// response map method of user
func (user *User) ResponseMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = user.ID
	res["email"] = user.Email
	res["first_name"] = user.FirstName
	res["last_name"] = user.LastName
	res["is_active"] = user.IsActive
	res["created_at"] = user.CreatedAt
	res["updated_at"] = user.UpdatedAt
	return res
}
