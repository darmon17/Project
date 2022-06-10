package resources

import (
	_userController "be9/project/controllers/user"
	_entities "be9/project/entities"

	"database/sql"
	"fmt"
)

func Register(db *sql.DB) int {
	newData := _entities.User{}
	fmt.Println("================")
	fmt.Print("Input Nama : ")
	fmt.Scanln(&newData.Nama)
	fmt.Print("Input Gender : ")
	fmt.Scanln(&newData.Gender)
	fmt.Print("Input Telepon : ")
	fmt.Scanln(&newData.Telp)
	fmt.Print("Input Password : ")
	fmt.Scanln(&newData.Password)
	fmt.Println("================")

	row, err := _userController.Create(db, newData)

	if err != nil {
		fmt.Println("gagal input", err.Error())
	}
	return row
}

func OtherProfile(db *sql.DB) []_entities.User {
	var phone string
	fmt.Println("Input others user phone: ")
	fmt.Scanln(&phone)
	results, err := _userController.OtherUser(db, phone)

	if err != nil {
		fmt.Println("error get data user", err)
	}
	return results
}

func Login(db *sql.DB) (_entities.User, error) {
	var datalogin _entities.User
	fmt.Println("================")
	fmt.Println("Form Login User")
	fmt.Println("================")
	fmt.Print("No telephone : ")
	fmt.Scanln(&datalogin.Telp)
	fmt.Print("Password : ")
	fmt.Scanln(&datalogin.Password)
	fmt.Println("================")
	result, err := _userController.Login(db, datalogin.Telp, datalogin.Password)
	if err != nil {
		panic(err.Error())
	} else {
		return result, nil
	}
}
