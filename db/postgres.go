package db

import (
	"database/sql"
	"log"
)

var (
	err error
)

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb(db *sql.DB) *PostgresDb {
	return &PostgresDb{
		db: db,
	}
}

func (repo PostgresDb) Migrate() error {
	log.Println("Migrating hotels")
	_, err = repo.db.Exec(CreateHotel)
	if err != nil {
		return err
	}
	log.Println("Finished migrating hotels")

	log.Println("Migrating floors")
	_, err = repo.db.Exec(CreateFloor)
	if err != nil {
		return err
	}
	log.Println("Finished migrating floors")

	log.Println("Migrating Rooms")
	_, err = repo.db.Exec(CreateRoom)
	if err != nil {
		return err
	}
	log.Println("Finished migrating rooms")

	log.Println("Migrating Beds")
	_, err = repo.db.Exec(CreateBed)
	if err != nil {
		return err
	}
	log.Println("Finished migrating Beds")

	return nil
}
