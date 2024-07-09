package models

import "github.com/google/uuid"

import "time"

type User struct {
	Id          uuid.UUID `json:"id" db:"id"`
	FirsrName   string    `json:"FirsrName" db:"FirsrName" `
	LastName    string    `json:"LastName" db:"LastName"`
	Email       string    `json:"Email" db:"Email"`
	Password    string    `json:"Password" db:"Password"`
	CreatedDate time.Time `json:"CreatedDate" db:"CreatedDate"`
	UpdatedDate time.Time `json:"UpdatedDate" db:"UpdatedDate"`
	IsDelete    bool      `json:"IsDelete" db:"IsDelete"`
}
