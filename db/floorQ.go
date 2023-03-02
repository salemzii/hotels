package db

import (
	"context"
	"hotels/models"
	"log"
)

// Create a new floor record
func (repo PostgresDb) CreateFloor(floor *models.Floor) (*models.Floor, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertFloor)
	if err != nil {
		return &models.Floor{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), floor.Floor_Number, floor.Hotel_id)
	var fl models.Floor
	e := row.Scan(&fl.Id, &fl.Floor_Number, &fl.Hotel_id)
	if e != nil {
		log.Println(e)
		return &models.Floor{}, e
	}
	return &fl, nil
}

// Fetch all floors records
func (repo PostgresDb) FetchAllFloor() (*[]models.Floor, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchAllFloors)
	if err != nil {
		return &[]models.Floor{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return &[]models.Floor{}, err
	}

	var floors []models.Floor

	for rows.Next() {
		var floor models.Floor
		if err = rows.Scan(&floor.Id, &floor.Floor_Number, &floor.Hotel_id); err != nil {
			log.Fatal(err)
		}
		floors = append(floors, floor)
	}

	return &floors, nil
}

// Fetch a floor record by id
func (repo PostgresDb) FetchFloor(floorId int) (*models.Floor, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchFloorById)
	if err != nil {
		return &models.Floor{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), floorId)
	var floor models.Floor

	err = row.Scan(&floor.Id, &floor.Floor_Number, &floor.Hotel_id)
	if err != nil {
		return &models.Floor{}, err
	}

	return &floor, nil
}

// Fetch a floor record by hotel id
func (repo PostgresDb) FetchFloorsByHotel(hotelId int) (*[]models.Floor, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchFloorByHotelId)
	if err != nil {
		return &[]models.Floor{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), hotelId)
	if err != nil {
		return &[]models.Floor{}, err
	}

	var floors []models.Floor

	for rows.Next() {
		var floor models.Floor
		if err = rows.Scan(&floor.Id, &floor.Floor_Number, &floor.Hotel_id); err != nil {
			log.Fatal(err)
		}
		floors = append(floors, floor)
	}

	return &floors, nil
}

func (repo PostgresDb) UpdateFloor(floor *models.Floor) (models.Floor, error) {
	return models.Floor{}, nil
}
