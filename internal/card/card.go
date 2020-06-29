package card

import "time"

// Card holds the card information
type Card struct {
	ID             int       `jsonapi:"primary,cards"`
	MemberID       string    `json:"memberId" jsonapi:"relation,members"`
	DisplayName    string    `json:"displayName" jsonapi:"attr,displayName"`
	LastFour       string    `json:"number" jsonapi:"attr,number"`
	ExpirationDate time.Time `json:"expirationDate" jsonapi:"attr,expirationDate"`
	BlockedAt      time.Time `json:"blockedAt" jsonapi:"attr,blockedAt"`
	CreatedAt      time.Time `json:"createdAt,omitempty" jsonapi:"attr,createdAt"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty" jsonapi:"attr,updatedAt"`
}
