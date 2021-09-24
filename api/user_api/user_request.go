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
	Password string `json:"password"`
}

type SetKeyUserRequest struct {
	OkexKey  OkexKeyDetail  `bson:"okex_key" json:"okex_key"`
	HuobiKey HuobiKeyDetail `bson:"huobi_key" json:"huobi_key"`
	PushUID  int64          `bson:"push_uid" json:"push_uid"`
}
