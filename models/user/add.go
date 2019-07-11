package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
	"time"
)

func AddNewUser(objUser dto.User) (newID int64, err error) {
	createdat := time.Now()
	rsl, err := driver.GetConn().Exec("INSERT INTO USER (UserName, FullName, CreatedUser, BirthDay, CreatedAt) VALUES (?,?,?,?,?)",
		objUser.UserName, objUser.FullName, objUser.CreatedUser, objUser.BirthDay, createdat)
	if err != nil {
		return -1, err
	}
	newID, err = rsl.LastInsertId()
	if err != nil {
		return -1, err
	}
	return
}
