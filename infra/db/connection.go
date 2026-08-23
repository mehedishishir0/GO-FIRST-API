package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection( conf config.DBConfig) string {
	
	conneionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name, conf.SSLMode)
	fmt.Println(conneionString)
	return conneionString
}

func NewConnection( conf config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnection(conf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil

}
