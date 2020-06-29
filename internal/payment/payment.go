package payment

import "time"

// Payment holds transaction's information.
// Relates member with a card in case of owning multiple cards.
type Payment struct {
	ID        int       `jsonapi:"primary,cards"`
	MemberID  string    `json:"memberId" jsonapi:"relation,members"`
	CardID    string    `json:"cardId" jsonapi:"relation,cards"`
	CreatedAt time.Time `json:"createdAt,omitempty" jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" jsonapi:"attr,updatedAt"`
}
