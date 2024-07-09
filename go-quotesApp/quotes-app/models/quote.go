package models

import (
	"time"

	"github.com/google/uuid"
)

type Quote struct {
	Id          uuid.UUID `json:"id" db:"id"`
	QuoteText   string    `json:"quote_text" db:"quote_text"`
	AuthorName  string    `json:"author_name" db:"author_name"`
	CreatedDate time.Time `json:"CreatedDate" db:"CreatedDate"`
	User        User
	UpdatedDate time.Time `json:"UpdatedDate" db:"UpdatedDate"`
	IsDelete    bool      `json:"IsDelete" db:"IsDelete"`
}
