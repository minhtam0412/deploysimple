package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
)

func GetUserByKey(id string) (user dto.User, err error) {
	cmd := "SELECT ID, UserName, FullName, CreatedUser, BirthDay, CreatedAt, UpdatedAt, UpdatedUser from user where ID = " + id
	err = driver.GetConn().Get(&user, cmd)
	if err != nil {
		return user, err
	}
	return
}
