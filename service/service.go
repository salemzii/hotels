package service

import (
	"database/sql"
	"errors"
	"hotels/db"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	Conn               string
	HotelRepository    *db.PostgresDb
	ErrCannotConnectDb = errors.New("unable to connect to database")
	ErrDuplicate       = errors.New("record already exists")
	ErrNotExists       = errors.New("row not exists")
	ErrDeleteFailed    = errors.New("delete failed")

	JwtSecretKey = []byte(os.Getenv("JwtSecretKey"))
)

func init() {

	/*
		if err := loadEnv(); err != nil {
			panic(err)
		}
	*/
	Conn = os.Getenv("CONNECTION") //viper.GetString("CONNECTION"

	database, err := sql.Open("postgres", Conn)
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	database.SetMaxIdleConns(20)

	HotelRepository = db.NewPostgresDb(database)
	if err := HotelRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

}

func loadEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.SetConfigType("env")
	viper.AddConfigPath(dir)
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
