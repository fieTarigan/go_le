package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpInput struct {
	UserGoogleId      string    `json:"user_google_id,omitempty" bson:"user_google_id,omitempty"`
	Username          string    `json:"username" bson:"username" binding:"required"`
	UserFullname      string    `json:"user_fullname" bson:"user_fullname" binding:"required"`
	UserEmail         string    `json:"user_email" bson:"user_email" binding:"required"`
	UserPassword      string    `json:"user_password" bson:"user_password" binding:"required,min=8"`
	UserImage         string    `json:"user_image,omitempty" bson:"user_image,omitempty"`
	UserImagePpgoogle string    `json:"user_image_ppgoogle,omitempty" bson:"user_image_ppgoogle,omitempty"`
	UserPhone         string    `json:"user_phone,omitempty" bson:"user_phone,omitempty"`
	UserAlamat        string    `json:"user_alamat,omitempty" bson:"user_alamat,omitempty"`
	UserDateOfBirth   time.Time `json:"user_date_of_birth,omitempty" bson:"user_date_of_birth,omitempty"`
	UserGender        string    `json:"user_gender,omitempty" bson:"user_gender,omitempty"`
	UserBio           string    `json:"user_bio,omitempty" bson:"user_bio,omitempty"`
	UserActive        uint8     `json:"user_active" bson:"user_active"`
	UserRole          uint8     `json:"user_role" bson:"user_role"`
	UserIsComplete    uint8     `json:"user_is_complete" bson:"user_is_complete"`
	UserLoginAt       time.Time `json:"user_login_at" bson:"user_login_at"`
	UserCreatedAt     time.Time `json:"user_created_at" bson:"user_created_at"`
}

type SignInInput struct {
	UserEmail    string `json:"user_email" bson:"user_email" binding:"required"`
	UserPassword string `json:"user_password" bson:"user_password" binding:"required"`
}

type DBResponse struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	UserGoogleId      string             `json:"user_google_id" bson:"user_google_id"`
	Username          string             `json:"username" bson:"username" binding:"required"`
	UserFullname      string             `json:"user_fullname" bson:"user_fullname" binding:"required"`
	UserEmail         string             `json:"user_email" bson:"user_email" binding:"required"`
	UserPassword      string             `json:"user_password" bson:"user_password" binding:"required"`
	UserImage         string             `json:"user_image" bson:"user_image"`
	UserImagePpgoogle string             `json:"user_image_ppgoogle" bson:"user_image_ppgoogle"`
	UserPhone         string             `json:"user_phone" bson:"user_phone"`
	UserAlamat        string             `json:"user_alamat" bson:"user_alamat"`
	UserDateOfBirth   time.Time          `json:"user_date_of_birth" bson:"user_date_of_birth"`
	UserGender        string             `json:"user_gender" bson:"user_gender"`
	UserBio           string             `json:"user_bio" bson:"user_bio"`
	UserActive        uint8              `json:"user_active" bson:"user_active"`
	UserRole          uint8              `json:"user_role" bson:"user_role"`
	UserIsComplete    uint8              `json:"user_is_complete" bson:"user_is_complete"`
	UserLoginAt       time.Time          `json:"user_login_at" bson:"user_login_at"`
	UserCreatedAt     time.Time          `json:"user_created_at" bson:"user_created_at"`
}

type UserResponse struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username          string             `json:"username,omitempty" bson:"username,omitempty"`
	UserFullname      string             `json:"user_fullname,omitempty" bson:"user_fullname,omitempty"`
	UserEmail         string             `json:"user_email,omitempty" bson:"user_email,omitempty"`
	UserImage         string             `json:"user_image,omitempty" bson:"user_image,omitempty"`
	UserImagePpgoogle string             `json:"user_image_ppgoogle,omitempty" bson:"user_image_ppgoogle,omitempty"`
	UserBio           string             `json:"user_bio,omitempty" bson:"user_bio,omitempty"`
	UserCreatedAt     time.Time          `json:"user_created_at" bson:"user_created_at"`
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		UserFullname:      user.UserFullname,
		UserEmail:         user.UserEmail,
		UserImage:         user.UserImage,
		UserImagePpgoogle: user.UserImagePpgoogle,
		UserBio:           user.UserBio,
		UserCreatedAt:     user.UserCreatedAt,
	}
}
