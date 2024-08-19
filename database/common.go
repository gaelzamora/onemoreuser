package database

import (
	"database/sql"
	"fmt"
	"onemoreuser/models"
	"onemoreuser/secrets"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	_, err := secrets.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	fmt.Println(SecretModel)
	Db, err = sql.Open("mysql", ConvertToString(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conection Success in Database")
	return nil
}
func ConvertToString(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}