package models

// LoginModel is a struct that represents the login model
type LoginModelInput struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// UserModel is a struct that represents the user model
type UserModelInput struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	FullName string `validate:"required" json:"full_name"`
}

type UserModelResponse struct {
	Email    string `validate:"required"`
	FullName string `validate:"required"`
}

type UserModelDB struct {
	ID        string `bson:"_id"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	FullName  string `bson:"full_name"`
	CreatedAt int64  `bson:"created_at"`
}

// WorkspaceModel is a struct that represents the workspace model
type WorkspaceModelInput struct {
	Name          string `validate:"required" json:"name"`
	CreatorUserId string `validate:"required" json:"creator_user_id"`
	MaxUsers      uint64 `validate:"required" json:"max_users"`
	Subscription  string `validate:"required" json:"subscription"`
}

type WorkspaceModelResponse struct {
	Name          string `validate:"required"`
	CreatorUserId string `validate:"required"`
	Subscription  string `validate:"required"`
	MaxUsers      uint64 `validate:"required"`
}

type WorkspaceModelDB struct {
	ID            string `bson:"_id"`
	Name          string `bson:"name"`
	CreatorUserId string `bson:"creator_user_id"`
	Subscription  string `bson:"subscription"`
	MaxUsers      uint64 `bson:"max_users"`
	ExpiredAt     int64  `bson:"expired_at"`
	CreatedAt     int64  `bson:"created_at"`
}
