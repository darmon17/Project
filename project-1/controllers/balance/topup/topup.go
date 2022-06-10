package topup

import (
	_entities "be9/project/entities"
	"database/sql"
)

/*
1. topup input nomor dan input saldo
2. saldo yang di inputkan akan bertambah ke nomor yang di inputkan (saldo awal + saldo akhir)
3. ambil nilai saldo yang di inputkan
4. terus update dg jumlah yang
*/
func TopUp(db *sql.DB, id int, telp string, saldo_topup int) (int, error) {
	var queryTopup = (`INSERT INTO topup_detail (user_id, telp, saldo_topup) VALUES (?,?,?)`)
	statement, errPrepare := db.Prepare(queryTopup)
	if errPrepare != nil {
		return 0, errPrepare
	}
	_, err := statement.Exec(id, telp, saldo_topup)
	if err != nil {
		return 0, err
	} else {
		//kondisi benar
		// proses update saldo user
		var queryGetSaldo = (`SELECT saldo FROM users WHERE id= ?`) //mengambil saldo awal dari tabel user
		getsaldo := db.QueryRow(queryGetSaldo, id)
		var saldo_awal int
		errGetsaldo := getsaldo.Scan(&saldo_awal)
		var saldoAkhir = saldo_awal + saldo_topup
		if errGetsaldo != nil {
			return 0, errGetsaldo
		} else {
			var updateSaldo = (`UPDATE users SET saldo= ? WHERE id= ?`)
			updateStatement, errUpdate := db.Prepare(updateSaldo)
			if errUpdate != nil {
				return 0, errUpdate
			}
			resultUpdatesaldo, errUpdatesaldo := updateStatement.Exec(saldoAkhir, id)
			if errUpdatesaldo != nil {
				return 0, errUpdatesaldo
			} else {
				rows, errRows := resultUpdatesaldo.RowsAffected()
				if errRows != nil {
					return 0, errRows
				} else {
					return int(rows), nil
				}

			}
		}

		/*
			ambil data saldo
			update users set saldo=saldo_topup + saldo where id=id;
		*/

	}
}

// menambhakan saldo user, dimana saldo diambil dari tabel topup_detail
//history topup
func Detail(db *sql.DB, id int) ([]_entities.TopUp, error) {
	var query = (`select id, saldo_topup, created_at from topup_detail where user_id = ? order by created_at DESC`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		panic(errPrepare.Error())
	}
	result, err := statement.Query(id)
	if err != nil {
		panic(err.Error())
	}
	var data []_entities.TopUp
	for result.Next() {
		var topup _entities.TopUp
		err := result.Scan(&topup.ID, &topup.Saldo_topup, &topup.Created_at)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, topup)
	}
	return data, nil
}
