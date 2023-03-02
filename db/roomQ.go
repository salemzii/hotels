package db

import (
	"context"
	"hotels/models"
	"log"
)

func (repo PostgresDb) CreateRoom(room *models.Room) (*models.Room, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertRoom)
	if err != nil {
		return &models.Room{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), room.Room_Number, room.Floor_id)
	var rm models.Room
	e := row.Scan(&rm.Id, &rm.Room_Number, &rm.Floor_id)
	if e != nil {
		log.Println(e)
		return &models.Room{}, e
	}
	return &rm, nil
}

func (repo PostgresDb) GetAllRoom() (*[]models.Room, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchAllRooms)
	if err != nil {
		return &[]models.Room{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return &[]models.Room{}, err
	}

	var rooms []models.Room
	for rows.Next() {
		var rm models.Room
		if err = rows.Scan(&rm.Id, &rm.Room_Number, &rm.Floor_id); err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, rm)
	}

	return &rooms, nil
}

func (repo PostgresDb) GetRoom(roomId int) (*models.Room, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchRoomById)
	if err != nil {
		return &models.Room{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), roomId)
	var rm models.Room

	err = row.Scan(&rm.Id, &rm.Room_Number, &rm.Floor_id)
	if err != nil {
		return &models.Room{}, err
	}

	return &rm, nil
}

func (repo PostgresDb) FetchRoomsByFloor(floorId int) (*[]models.Room, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchRoomByFloorId)
	if err != nil {
		return &[]models.Room{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), floorId)
	if err != nil {
		return &[]models.Room{}, err
	}

	var rooms []models.Room
	for rows.Next() {
		var rm models.Room
		if err = rows.Scan(&rm.Id, &rm.Room_Number, &rm.Floor_id); err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, rm)
	}

	return &rooms, nil
}

func (repo PostgresDb) UpdateRoom(room *models.Room) (models.Room, error) {
	return models.Room{}, nil
}
