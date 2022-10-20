package domain

import (
	"github.com/ervera/tdlc-gin/pkg/iso8601"
	"github.com/google/uuid"
)

type Team struct {
	ID        int           `json:"id,omitempty"`
	UUID      uuid.UUID     `json:"uuid,omitempty"`
	CreatedOn iso8601.ITime `json:"created_on,omitempty"`
	Name      string        `json:"name,omitempty"`
	Enable    bool          `json:"enable,omitempty"`
	Color     string        `json:"color,omitempty"`
	Image     string        `json:"image,omitempty"`
	Role      Role          `json:"role,omitempty"`
}

type Role string

const (
	Owner  Role = "owner"
	Admin  Role = "admin"
	Member Role = "member"
)
