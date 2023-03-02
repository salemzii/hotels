package db

import (
	"context"
	"hotels/models"
	"log"
)

func (repo PostgresDb) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertHotel)
	if err != nil {
		return &models.Hotel{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), hotel.Name, hotel.Address, hotel.Prefix)

	var h models.Hotel
	err = row.Scan(&h.Id, &h.Name, &h.Address, &h.Prefix)
	if err != nil {
		return &models.Hotel{}, err
	}

	return &h, nil
}

func (repo PostgresDb) FetchAllHotel() (*[]models.Hotel, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchAllHotels)
	if err != nil {
		return &[]models.Hotel{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return &[]models.Hotel{}, err
	}

	var hotels []models.Hotel

	for rows.Next() {
		var h models.Hotel
		if err = rows.Scan(&h.Id, &h.Name, &h.Address, &h.Prefix); err != nil {
			log.Fatal(err)
		}
		hotels = append(hotels, h)
	}

	return &hotels, nil
}

func (repo PostgresDb) FetchHotel(id int) (*models.Hotel, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchHotelById)
	if err != nil {
		return &models.Hotel{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), id)
	var h models.Hotel

	err = row.Scan(&h.Id, &h.Name, &h.Address, &h.Prefix)
	if err != nil {
		return &models.Hotel{}, err
	}

	return &h, nil
}

func (repo PostgresDb) UpdateHotel(hotel *models.Hotel) (models.Hotel, error) {
	return models.Hotel{}, nil
}
