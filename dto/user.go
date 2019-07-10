package dto

import "time"

type User struct {
	ID          int       `db:"ID" json:"id"`
	UserName    string    `db:"UserName" json:"username"`
	FullName    string    `db:"FullName" json:"fullname"`
	CreatedDate time.Time `db:"CreatedDate" json:"createddate"`
	CreatedUser string    `db:"CreatedUser" json:"createduser"`
	UpdatedDate time.Time `db:"UpdatedDate" json:"updateddate"`
	UpdatedUser string    `db:"UpdatedUser" json:"updateduser"`
}
