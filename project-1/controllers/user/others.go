package user

import (
	_entities "be9/project/entities"

	"database/sql"
	"fmt"
)

func OtherUser(db *sql.DB, phone string) ([]_entities.User, error) {
	var query = (`SELECT nama, telp FROM users WHERE telp = ?`)
	statement, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	result, err := statement.Query(phone)
	if err != nil {
		panic(err.Error())
	}

	var data []_entities.User
	for result.Next() {
		var user _entities.User
		err := result.Scan(&user.Nama, &user.Telp)

		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		data = append(data, user)
	}
	return data, nil
}
