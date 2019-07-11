package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
)

func GetAllUser() (lstUser []dto.User, err error) {
	cmd := "SELECT ID, UserName,FullName,BirthDay, CreatedUser,CreatedAt,UpdatedUser,UpdatedAt FROM USER"
	driver.GetConn().Select(&lstUser, cmd)
	if err != nil {
		return nil, err
	}
	return
}
