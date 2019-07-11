package user

import (
	"deploysimple/driver"
	"deploysimple/dto"
	"time"
)

func UpdateUser(user dto.User) (err error) {
	updateat := time.Now()
	stmt, err := driver.GetConn().Prepare("UPDATE user set FullName=?, BirthDay=?, UpdatedUser=?, UpdatedAt=? where ID = ?")
	_, err = stmt.Exec(user.FullName, user.BirthDay, user.UpdatedUser, updateat, user.ID)
	stmt.Close()
	return
}
