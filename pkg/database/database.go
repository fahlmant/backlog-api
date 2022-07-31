package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)


var (
	DB *sql.DB
)

func InitDb () {
	var err error
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config["dbhost"], config["dbport"],
		config["dbuser"], config["dbpass"], config["dbname"])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	viper.AddConfigPath("../../pkg/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	
	conf := make(map[string]string)
	host := viper.GetString("database.dbhost")
	port := viper.GetString("database.dbport")
	user := viper.GetString("database.dbuser")
	password := viper.GetString("database.dbpass")
	name := viper.GetString("database.dbname")

	conf["dbhost"] = host
	conf["dbport"] = port
	conf["dbuser"] = user
	conf["dbpass"] = password
	conf["dbname"] = name
	return conf
	}