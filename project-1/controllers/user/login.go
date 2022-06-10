package user

import (
	_entities "be9/project/entities"

	"database/sql"
)

func Login(db *sql.DB, telp, password string) (_entities.User, error) {
	var user _entities.User
	if eror := db.QueryRow("SELECT id, nama, telp, saldo FROM users WHERE telp = ? AND password = ?", telp, password).Scan(&user.ID, &user.Nama, &user.Telp, &user.Saldo); eror != nil {
		if eror == sql.ErrNoRows {
			return _entities.User{}, eror
		}
		return _entities.User{}, eror
	}
	return user, nil
}
