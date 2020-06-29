package api

import "time"

// Member ...
type Member struct {
	ID        int       `jsonapi:"primary,members"`
	Name      string    `json:"name" jsonapi:"attr,name" validate:"nonzero,max=40,regexp=^[a-zA-Z]+( [a-zA-Z]+)*$"`
	Email     string    `json:"email" jsonapi:"attr,email" validate:"nonzero,max=40,regexp=^[a-zA-Z0-9_+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-]+$"`
	Phone     string    `json:"phone" jsonapi:"attr,phone" validate:"min=14,max=14"`
	Password  string    `json:"password" jsonapi:"attr,password" validate:"min=8"`
	CreatedAt time.Time `json:"createdAt,omitempty" jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" jsonapi:"attr,updatedAt"`
}
