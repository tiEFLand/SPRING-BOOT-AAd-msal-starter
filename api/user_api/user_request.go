package user_api

// Create User
// we get uid from the jwt
type CreateUserRequest struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type DeleteUserRequest struct {
	UID string `json:"uid"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Pass