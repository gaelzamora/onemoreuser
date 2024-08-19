package database

import (
	"fmt"
	"onemoreuser/models"
	"onemoreuser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start Register")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "INSERT INTO User (User_Email, User_Id, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateMySQL() + "')"

	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Signup > Successful")
	return nil
}