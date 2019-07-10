package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
)

func GetAllUser() (lstUser []dto.NewUser, err error) {
	cmd := "select ID, UserName, FullName, CreatedUser, BirthDay, CreatedAt, UpdatedAt, UpdatedUser FROM user"
	driver.GetConn().Select(&lstUser, cmd)
	if err != nil {
		return nil, err
	}
	return
}
