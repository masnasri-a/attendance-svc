package models

type AttendanceModelDB struct {
	ID             string `bson:"_id"`
	UserId         string `bson:"user_id"`
	AttendanceType string `bson:"attendance_type"`
	Longitute      string `bson:"longitute"`
	Latitude       string `bson:"latitude"`
	CreatedAt      int64  `bson:"created_at"`
}

type AttendanceModelInput struct {
	AttendanceType string `validate:"required" json:"attendance_type"`
	Longitute      string `validate:"required" json:"longitute"`
	Latitude       string `validate:"required" json:"latitude"`
}