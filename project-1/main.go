package main

import (
	_config "be9/project/config"
	_topupController "be9/project/controllers/balance/topup"
	_transferController "be9/project/controllers/balance/transfer"
	_userController "be9/project/controllers/user"
	_menuResources "be9/project/resources"

	_entities "be9/project/entities"

	"database/sql"
	"fmt"
)

var DBConn *sql.DB

func init() {
	DBConn = _config.ConnectionDB()
}

func main() {
	defer DBConn.Close()
	fmt.Println("1: Register")
	fmt.Println("2: Login")
	fmt.Println("================")
	fmt.Print("Pilih Menu : ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		rows := _menuResources.Register(DBConn)

		fmt.Println("Registrasi Success")
		fmt.Println("Status OK", rows)
		fmt.Println("Next Login ?")
		fmt.Println("(1: Login)/(2: Logout)")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			check, err := _menuResources.Login(DBConn)
			if err != nil {
				fmt.Println("Login Failed! cek no telp & pasword", check.Gender)
			} else {
				fmt.Println("Login Succes! Hallo,", check.Nama)
				fmt.Println("================")
				//menu setelah success atau menu utama
				fmt.Println("Pilih Menu Utama")
				fmt.Println("================")
				fmt.Println("1: Profil")
				fmt.Println("2: Transaksi")
				fmt.Println("3: Lihat Profil Teman")
				fmt.Println("0: Logout")
				fmt.Println("================")
				fmt.Print("Pilih Menu : ")
				var menu int
				fmt.Scanln(&menu)

				switch menu {
				case 1:
					fmt.Println("***** ", check.Nama, " *****")
					fmt.Println("Telepon:", check.Telp, "Saldo:", check.Saldo)
					var pilihan int
					fmt.Println("1: Update")
					fmt.Println("2: Delete Account")
					fmt.Println("================")
					fmt.Print("Pilih Menu : ")
					fmt.Scanln(&pilihan)

					if pilihan == 1 {
						newData := _entities.User{}
						var id int = check.ID

						fmt.Println("Input Nama : ")
						fmt.Scan(&newData.Nama)
						fmt.Println("================")

						_, err := _userController.Update(DBConn, newData, id)

						if err != nil {
							fmt.Println("gagal update", err.Error())
						} else {
							fmt.Println("Berhasil update Data!")
						}
					}
					if pilihan == 2 {
						var id int = check.ID
						fmt.Println("================")
						fmt.Println("Yakin ingin di hapus ?")
						fmt.Println("1: OK")
						fmt.Println("2: Cancel")
						fmt.Println("================")
						fmt.Print("Pilih Menu : ")
						var pilih int
						fmt.Scanln(&pilih)
						if pilih == 1 {
							_, err := _userController.Delete(DBConn, id)
							if err != nil {
								fmt.Println("gagal hapus data", err.Error())
							} else {
								fmt.Println("Berhasil hapus data!")
							}
						} else {
							fmt.Println("***** Kamu yang terbaik *****")
						}
					}
				case 2: //transaksi
					fmt.Println("================")
					fmt.Println("Silahkan pilih menu")
					fmt.Println("Transaksi : ")
					fmt.Println("1) Transfer ")
					fmt.Println("2) Top-up")
					fmt.Println("3) History top-up ")
					fmt.Println("4) History transfer ")
					fmt.Println("5) Logout")
					fmt.Println("=============")
					fmt.Print("Pilih menu : ")
					var chooseTrx int
					fmt.Scan(&chooseTrx)

					if chooseTrx == 1 {
						fmt.Println("ini halaman transfer")
						//transfer code
						var telp_receiver string
						var saldo_transfer int
						fmt.Println("Telepon Penerima")
						fmt.Scanln(&telp_receiver)
						fmt.Println("Jumlah Saldo Kirim")
						fmt.Scanln(&saldo_transfer)
						row, err := _transferController.Transfer(DBConn, check.ID, saldo_transfer, check.Telp, telp_receiver)
						if err != nil {
							panic(err.Error())
						} else if row != 0 {
							fmt.Println("Transfer Berhasil")
							fmt.Println("row affect", row)
						} else {
							fmt.Println("Saldo anda kurang !")
							fmt.Println("row affect", row)
						}
					}

					if chooseTrx == 2 {
						fmt.Println("================")
						//topup code
						var saldo int
						fmt.Println("Nomor Telphone ", check.Telp)
						fmt.Print("Input Nominal : ")
						fmt.Scan(&saldo)
						_, err := _topupController.TopUp(DBConn, check.ID, check.Telp, saldo)
						if err != nil {
							fmt.Println("Eror")
						} else {
							fmt.Println("================")
							fmt.Println("Top-up Succes!")
						}

					}
					if chooseTrx == 3 {
						fmt.Println("=================================================================")
						fmt.Println("Hallo ", check.Nama, "Berikut history top up anda")
						detailTopup, detailErr := _topupController.Detail(DBConn, check.ID)
						if detailErr != nil {
							panic(detailErr.Error())
						}
						for _, value := range detailTopup {
							fmt.Print("ID Transaksi : ", value.ID, " ")
							fmt.Print("Saldo Topup : Rp ", value.Saldo_topup, " Tanggal Transaksi ")
							for _, v := range value.Created_at {
								fmt.Print(string(v))
							}
							fmt.Print("\n")
						}
						fmt.Println("=================================================================")
					}
					if chooseTrx == 4 {
						checkHistory, errHistory := _transferController.History(DBConn, check.ID)
						if errHistory != nil {
							panic(errHistory.Error())
						} else {
							for _, value := range checkHistory {
								fmt.Println("Nama: ", value.Nama)
								fmt.Println("Pengirim: ", value.TelpSender)
								fmt.Println("Penerima: ", value.TelpReceiver)
								fmt.Print("Saldo Transfer: ", value.SaldoTransfer, "  date: ")
								for _, v := range value.CreateAt {
									fmt.Print(string(v))
								}
								fmt.Print("\n")
							}
						}
					}

				case 3:
					others := _menuResources.OtherProfile(DBConn)
					for _, v := range others {
						fmt.Println("================")
						fmt.Println("Nama:", v.Nama)
						fmt.Println("Telepon:", v.Telp)
						fmt.Println("================")
					}

				case 0:
					fmt.Println("Terima Kasih Telah Bertransaksi")
				}
			}
		}

	case 2:
		check, err := _menuResources.Login(DBConn)
		if err != nil {
			fmt.Println("Login Failed! cek no telp & pasword")
		} else {
			fmt.Println("Login Succes! Hallo,", check.Nama)
			fmt.Println("================")
			//menu setelah success atau menu utama
			fmt.Println("Pilih Menu Utama")
			fmt.Println("================")
			fmt.Println("1: Profil")
			fmt.Println("2: Transaksi")
			fmt.Println("3: Lihat Profil Teman")
			fmt.Println("0: Logout")
			fmt.Println("================")
			fmt.Print("Pilih Menu : ")
			var menu int
			fmt.Scanln(&menu)

			switch menu {
			case 1:
				fmt.Println("***** ", check.Nama, " *****")
				fmt.Println("Telepon:", check.Telp, "Saldo:", check.Saldo)
				var pilihan int
				fmt.Println("1: Update")
				fmt.Println("2: Delete Account")
				fmt.Println("================")
				fmt.Print("Pilih Menu : ")
				fmt.Scanln(&pilihan)

				if pilihan == 1 {
					newData := _entities.User{}
					var id int = check.ID

					fmt.Print("Input Nama : ")
					fmt.Println("================")
					fmt.Scanln(&newData.Nama)

					_, err := _userController.Update(DBConn, newData, id)

					if err != nil {
						fmt.Println("gagal update", err.Error())
					} else {
						fmt.Println("Berhasil update Data!")
					}
				}
				if pilihan == 2 {
					var id int = check.ID
					fmt.Println("================")
					fmt.Println("Yakin ingin di hapus ?")
					fmt.Println("1: OK")
					fmt.Println("2: Cancel")
					fmt.Println("================")
					fmt.Print("Pilih Menu : ")
					var pilih int
					fmt.Scanln(&pilih)
					if pilih == 1 {
						_, err := _userController.Delete(DBConn, id)
						if err != nil {
							fmt.Println("gagal hapus data", err.Error())
						} else {
							fmt.Println("Berhasil hapus data!")
						}
					} else {
						fmt.Println("***** Kamu yang terbaik *****")
					}
				}
			case 2: //transaksi
				fmt.Println("================")
				fmt.Println("Silahkan pilih menu")
				fmt.Println("Transaksi : ")
				fmt.Println("1) Transfer ")
				fmt.Println("2) Top-up")
				fmt.Println("3) History top-up ")
				fmt.Println("4) History transfer ")
				fmt.Println("5) Logout")
				fmt.Println("=============")
				fmt.Print("Pilih menu : ")
				var chooseTrx int
				fmt.Scan(&chooseTrx)

				if chooseTrx == 1 {
					fmt.Println("ini halaman transfer")
					//transfer code
					var telp_receiver string
					var saldo_transfer int
					fmt.Println("Telepon Penerima")
					fmt.Scanln(&telp_receiver)
					fmt.Println("Jumlah Saldo Kirim")
					fmt.Scanln(&saldo_transfer)
					row, err := _transferController.Transfer(DBConn, check.ID, saldo_transfer, check.Telp, telp_receiver)
					if err != nil {
						panic(err.Error())
					} else if row != 0 {
						fmt.Println("Transfer Berhasil")
						fmt.Println("row affect", row)
					} else {
						fmt.Println("Saldo anda kurang !")
						fmt.Println("row affect", row)
					}
				}

				if chooseTrx == 2 {
					fmt.Println("================")
					//topup code
					var saldo int
					fmt.Println("Nomor Telphone ", check.Telp)
					fmt.Print("Input Nominal : ")
					fmt.Scan(&saldo)
					_, err := _topupController.TopUp(DBConn, check.ID, check.Telp, saldo)
					if err != nil {
						fmt.Println("Eror")
					} else {
						fmt.Println("================")
						fmt.Println("Top-up Succes!")
					}

				}
				if chooseTrx == 3 {
					fmt.Println("=================================================================")
					fmt.Println("Hallo ", check.Nama, "Berikut history top up anda")
					detailTopup, detailErr := _topupController.Detail(DBConn, check.ID)
					if detailErr != nil {
						panic(detailErr.Error())
					}
					for _, value := range detailTopup {
						fmt.Print("ID Transaksi : ", value.ID, " ")
						fmt.Print("Saldo Topup : Rp ", value.Saldo_topup, " Tanggal Transaksi ")
						for _, v := range value.Created_at {
							fmt.Print(string(v))
						}
						fmt.Print("\n")
					}
					fmt.Println("=================================================================")
				}
				if chooseTrx == 4 {
					checkHistory, errHistory := _transferController.History(DBConn, check.ID)
					if errHistory != nil {
						panic(errHistory.Error())
					} else {
						for _, value := range checkHistory {
							fmt.Println("Nama: ", value.Nama)
							fmt.Println("Pengirim: ", value.TelpSender)
							fmt.Println("Penerima: ", value.TelpReceiver)
							fmt.Print("Saldo Transfer: ", value.SaldoTransfer, "  date: ")
							for _, v := range value.CreateAt {
								fmt.Print(string(v))
							}
							fmt.Print("\n")
						}
					}
				}

			case 3:
				others := _menuResources.OtherProfile(DBConn)
				for _, v := range others {
					fmt.Println("================")
					fmt.Println("Nama:", v.Nama)
					fmt.Println("Telepon:", v.Telp)
					fmt.Println("================")
				}

			case 0:
				fmt.Println("Terima Kasih Telah Bertransaksi")
			}
		}
	}
}
