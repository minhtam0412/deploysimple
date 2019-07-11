package user

import (
	"deploysimple/driver"
	"time"
)

func DeleteUser(id string, deleteduser string) (err error) {
	deletedat := time.Now()
	stmt, err := driver.GetConn().Prepare("UPDATE user set IsDeleted = 1, DeletedAt=?, DeletedUser=? where  ID = ?")
	_, err = stmt.Exec(deletedat, deleteduser, id)
	stmt.Close()
	return
}
