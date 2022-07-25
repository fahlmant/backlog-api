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

func dbInit() {
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
	viper.AddConfigPath("/Users/jordanlange/Documents/projects/personal/backlog-api/pkg/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	
	conf := make(map[string]string)
	host := viper.GetString("database.dbhost")
	fmt.Println(host)
	port := viper.GetString("database.dbport")
	fmt.Println(port)
	user := viper.GetString("database.dbuser")
	fmt.Println(user)
	password := viper.GetString("database.dbpass")
	fmt.Println(password)
	name := viper.GetString("database.dbname")
	fmt.Println(name)

	conf["dbhost"] = host
	conf["dbport"] = port
	conf["dbuser"] = user
	conf["dbpass"] = password
	conf["dbname"] = name
	fmt.Println(conf)
	return conf
	}