package dto

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int            `db:"ID" json:"id"`
	UserName    string         `db:"UserName" json:"username,omitempty"`
	FullName    string         `db:"FullName" json:"fullname,omitempty"`
	BirthDay    mysql.NullTime `db:"BirthDay" json:"birthday,omitempty"`
	CreateAt    mysql.NullTime `db:"CreatedAt" json:"createdat,omitempty"`
	CreatedUser sql.NullString `db:"CreatedUser" json:"createduser"`
	UpdatedAt   mysql.NullTime `db:"UpdatedAt" json:"updatedat"`
	UpdatedUser sql.NullString `db:"UpdatedUser" json:"updateduser"`
	IsDeleted   bool           `db:"IsDeleted" json:"isdeleted"`
	DeletedAt   mysql.NullTime `db:"DeletedAt" json:"deletedat"`
	DeletedUser sql.NullString `db:"DeletedUser" json:"deleteduser"`
}
