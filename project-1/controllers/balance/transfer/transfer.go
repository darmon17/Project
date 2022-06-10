package transfer

import (
	"database/sql"
)

func GetSaldo(db *sql.DB, telp string) int {
	var saldo int
	var query = (`SELECT saldo FROM users WHERE telp = ?`)
	err := db.QueryRow(query, telp)
	err.Scan(&saldo)
	return saldo
}

func UpdateSaldoSender(db *sql.DB, saldo int, telp string) (int, error) {
	var updateSaldoSender = (`UPDATE users SET saldo = ? WHERE telp = ?`)
	updateStatement, errUpdate := db.Prepare(updateSaldoSender)
	if errUpdate != nil {
		return 0, errUpdate
	}
	result, errSaldoSender := updateStatement.Exec(saldo, telp)
	if errSaldoSender != nil {
		return 0, nil
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func UpdateSaldoReceiver(db *sql.DB, saldo int, telp string) (int, error) {
	var updateSaldo = (`UPDATE users SET saldo = ? WHERE telp = ?`)
	updateStatement, errUpdate := db.Prepare(updateSaldo)
	if errUpdate != nil {
		return 0, errUpdate
	}
	result, errSaldoSender := updateStatement.Exec(saldo, telp)
	if errSaldoSender != nil {
		return 0, nil
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func Transfer(db *sql.DB, id, saldo_transfer int, telp_sender, telp_receiver string) (int, error) {
	saldoSender := GetSaldo(db, telp_sender)
	saldoReceiver := GetSaldo(db, telp_receiver)

	newSaldoSender := saldoSender - saldo_transfer
	newSaldoReceiver := saldoReceiver + saldo_transfer

	if saldoSender < saldo_transfer {
		return 0, nil
	} else {
		var query = (`INSERT INTO transfer_detail (user_id, telp_sender, telp_receiver, saldo_transfer, status_id) VALUES (?,?,?,?,?)`)
		statement, errPrepare := db.Prepare(query)

		if errPrepare != nil {
			return 0, errPrepare
		}
		_, err := statement.Exec(&id, &telp_sender, &telp_receiver, &saldo_transfer, 0)

		if err != nil {
			return 0, err
		}
		_, errSender := UpdateSaldoSender(db, newSaldoSender, telp_sender)
		if errSender != nil {
			return 0, errSender
		} else {
			result, errReceiver := UpdateSaldoReceiver(db, newSaldoReceiver, telp_receiver)
			if errReceiver != nil {
				return 0, errReceiver
			} else {
				return result, nil
			}
		}
	}
}

// func Status(db *sql.DB, telp string) int {
// 	var status int
// 	var query = (`SELECT status_id FROM transfer_detail WHERE telp_sender = ?`)
// 	err := db.QueryRow(query, telp)
// 	err.Scan(&status)
// 	return status
// }
