package transfer

import (
	_entities "be9/project/entities"

	"database/sql"
	"fmt"
)

func History(db *sql.DB, id int) ([]_entities.Transfer, error) {
	var query = (`SELECT u.nama, trd.telp_sender, trd.telp_receiver, trd.saldo_transfer, trd.created_at FROM users u INNER JOIN transfer_detail trd ON u.id = trd.user_id WHERE u.id = ? ORDER BY created_at DESC `)

	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		panic(errPrepare.Error())
	}

	result, err := statement.Query(id)
	if err != nil {
		panic(err.Error())
	}

	var data []_entities.Transfer
	for result.Next() {
		var user _entities.Transfer
		err := result.Scan(&user.Nama, &user.TelpSender, &user.TelpReceiver, &user.SaldoTransfer, &user.CreateAt)

		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		data = append(data, user)
	}
	return data, nil
}
