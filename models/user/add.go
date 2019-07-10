package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
)

func AddNewUser(objUser dto.NewUser) (newID int64, err error) {
	rsl, err := driver.GetConn().Exec("Insert into user (UserName, FullName, CreatedAt, CreatedUser, BirthDay, UpdatedAt, UpdatedUser) values (?,?,?,?,?,?,?)",
		objUser.UserName, objUser.FullName, objUser.CreateAt, objUser.CreatedUser, objUser.BirthDay, objUser.UpdateAt, objUser.UpdatedUser)
	if err != nil {
		return -1, err
	}
	newID, err = rsl.LastInsertId()
	if err != nil {
		return -1, err
	}
	return
}
