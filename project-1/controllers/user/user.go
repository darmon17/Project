package user

import (
	_entities "be9/project/entities"

	"database/sql"
)

func Create(db *sql.DB, newTopup _entities.User) (int, error) {
	var query = (`INSERT INTO users (nama, gender, telp, password, saldo) VALUES (?,?,?,?,?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(newTopup.Nama, newTopup.Gender, newTopup.Telp, newTopup.Password, 0)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func Update(db *sql.DB, newData _entities.User, id int) (int, error) {
	var query = (`UPDATE users SET nama = ? WHERE id = ? `)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	defer db.Close()

	result, err := statement.Exec(newData.Nama, id)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func Delete(db *sql.DB, id int) (int, error) {
	var query = (`DELETE FROM users WHERE id = ? `)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	defer db.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
