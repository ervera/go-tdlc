package domain

import (
	"github.com/ervera/tdlc-gin/pkg/iso8601"
	"github.com/google/uuid"
)

type User struct {
	//ID        int64         `json:"id,omitempty"`
	ID        uuid.UUID     `json:"id,omitempty"`
	FirstName string        `json:"first_name,omitempty"`
	LastName  string        `json:"last_name,omitempty"`
	Password  string        `json:"password,omitempty"`
	Email     string        `json:"email,omitempty" `
	CreatedOn iso8601.ITime `json:"created_on,omitempty"`
	Avatar    string        `json:"avatar,omitempty"`
	Banner    string        `json:"banner,omitempty"`
	Biography string        `json:"biography,omitempty"`
	Location  string        `json:"location,omitempty"`
	Website   string        `json:"website,omitempty"`
	Token     string        `json:"token,omitempty"`
}
