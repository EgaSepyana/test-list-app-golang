package model

import (
	"time"
	"todolist-api/src/util/umongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generated by https://quicktype.io

type User struct {
	MetadataWithID `bson:",inline"`
	Email          *string `json:"email" bson:"email,omitempty"`
	Username       string  `json:"username" bson:"username,omitempty"`
	Password       string  `json:"password" bson:"password,omitempty"`
}

type UserRegister struct {
	MetadataWithID `bson:",inline"`

	Email    string `json:"email,omitempty" bson:"email"`
	UserName string `json:"username,omitempty" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
}

type User_View struct {
	User `bson:",inline"`
}

type User_Search struct {
	//? Regex
	Search string `json:"search"`

	umongo.Request
}

func (o *User_Search) HandleFilter(listFilterAnd *[]bson.M) {
	if search := o.Search; search != "" {
		*listFilterAnd = append(*listFilterAnd, bson.M{"username": primitive.Regex{Pattern: search, Options: "i"}})
	}
}

type User_Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	//? SSO
	Email string `json:"email,omitempty" swaggerignore:"true"`
	AppId string `json:"appId,omitempty" swaggerignore:"true"`
	OtpId string `json:"otpId,omitempty" swaggerignore:"true"`
}

type User_ChangePassword struct {
	Id          string `json:"-" bson:"_id,omitempty" swaggerignore:"true"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type Resp_JwtToken struct {
	Token  string    `json:"token,omitempty"`
	Expire time.Time `json:"expire,omitempty"`
}

type Resp_User_Login struct {
	User *User         `json:"user,omitempty"`
	Auth Resp_JwtToken `json:"auth,omitempty"`
}

// Generated by https://quicktype.io

type EditProfile struct {
	IdDocument string `json:"-" bson:"_id,omitempty" swaggerignore:"true"`

	Avatar      string `json:"avatar" bson:"avatar,omitempty"`
	Fullname    string `json:"fullname" bson:"fullname,omitempty"`
	Username    string `json:"username" bson:"username,omitempty"`
	About       string `json:"about" bson:"about,omitempty"`
	Country     string `json:"country" bson:"country,omitempty"`
	Email       string `json:"email" bson:"email,omitempty"`
	Phone       string `json:"phone" bson:"phone,omitempty"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
}
