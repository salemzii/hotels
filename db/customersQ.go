package db

import (
	"context"
	"errors"
	"hotels/models"
)

var (
	ErrFetchingBed = errors.New("unable to fetch bed")
	ErrOccupiedBed = errors.New("bed is occupied")
	ErrCheckoutBed = errors.New("bed is on checkout")
)

func (repo PostgresDb) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertCustomer)
	if err != nil {
		return &models.Customer{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), customer.Name, customer.Agent_name)

	var c models.Customer
	err = row.Scan(&c.Id, &c.Name, &c.Agent_name)
	if err != nil {
		return &models.Customer{}, err
	}

	return &c, nil
}

func (repo PostgresDb) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertCustomerReservation)
	if err != nil {
		return &models.Reservation{}, err
	}
	defer stmt.Close()

	reservedBed, err := repo.FetchBed(reservation.Bed_id)
	if err != nil {
		return &models.Reservation{}, ErrFetchingBed
	}

	switch reservedBed.Status {
	case "occupied":
		return &models.Reservation{}, ErrOccupiedBed
	case "checkout":
		return &models.Reservation{}, ErrCheckoutBed
	}

	row := stmt.QueryRowContext(context.Background(), reservation.From_date, reservation.To_date, reservation.Bed_id, reservation.Customer_id)
	var rsv models.Reservation
	err = row.Scan(&rsv.Id, &rsv.From_date, &rsv.To_date, &rsv.Bed_id, &rsv.Customer_id)
	if err != nil {
		return &models.Reservation{}, err
	}

	reservedBed.Status = "occupied"
	_, err = repo.UpdateBedStatus(reservedBed)
	if err != nil {
		return &rsv, err
	}
	return &rsv, nil
}
