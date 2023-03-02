package db

import (
	"context"
	"hotels/models"
	"log"
)

func (repo PostgresDb) CreateBed(bed *models.Bed) (*models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), InsertBed)
	if err != nil {
		return &models.Bed{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), bed.Bed_Number, bed.Price, bed.Status, bed.Room_id)
	var bd models.Bed
	err = row.Scan(&bd.Id, &bd.Bed_Number, &bd.Price, &bd.Status, &bd.Room_id)
	if err != nil {
		return &models.Bed{}, err
	}
	return &bd, nil
}

func (repo PostgresDb) UpdateBedStatus(bed *models.Bed) (*models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), UpdateBedStatus)
	if err != nil {
		return &models.Bed{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), bed.Room_id, bed.Status)

	var bd models.Bed
	err = row.Scan(&bd.Id, &bd.Bed_Number, &bd.Price, &bd.Status, &bd.Room_id)
	if err != nil {
		return &models.Bed{}, err
	}
	return &bd, nil
}

func (repo PostgresDb) FetchAllBed() (*[]models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchAllBeds)
	if err != nil {
		return &[]models.Bed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return &[]models.Bed{}, err
	}

	var beds []models.Bed

	for rows.Next() {
		var bed models.Bed
		if err = rows.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id); err != nil {
			log.Fatal(err)
		}
		beds = append(beds, bed)
	}

	return &beds, nil
}

func (repo PostgresDb) FetchBed(bedId int) (*models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchBedById)
	if err != nil {
		return &models.Bed{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), bedId)
	var bed models.Bed

	err = row.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id)
	if err != nil {
		return &models.Bed{}, err
	}

	return &bed, nil
}

func (repo PostgresDb) FetchBedsByRoom(roomId int) (*[]models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchBedByRoomId)
	if err != nil {
		return &[]models.Bed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), roomId)
	if err != nil {
		return &[]models.Bed{}, err
	}

	var beds []models.Bed

	for rows.Next() {
		var bed models.Bed
		if err = rows.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id); err != nil {
			log.Fatal(err)
		}
		beds = append(beds, bed)
	}

	return &beds, nil
}

func (repo PostgresDb) FetchAvailableBeds(roomId int)(*[]models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchAvailableBedsByRoomId)
	if err != nil {
		return &[]models.Bed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), roomId)
	if err != nil {
		return &[]models.Bed{}, err
	}

	var beds []models.Bed

	for rows.Next() {
		var bed models.Bed
		if err = rows.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id); err != nil {
			log.Fatal(err)
		}
		beds = append(beds, bed)
	}

	return &beds, nil
}

func (repo PostgresDb) FetchOccupiedBeds(roomId int)(*[]models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchOccupiedBedsByRoomId)
	if err != nil {
		return &[]models.Bed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), roomId)
	if err != nil {
		return &[]models.Bed{}, err
	}

	var beds []models.Bed

	for rows.Next() {
		var bed models.Bed
		if err = rows.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id); err != nil {
			log.Fatal(err)
		}
		beds = append(beds, bed)
	}

	return &beds, nil
}

func (repo PostgresDb) FetchCheckOutBeds(roomId int)(*[]models.Bed, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), FetchCheckoutBedsByRoomId)
	if err != nil {
		return &[]models.Bed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), roomId)
	if err != nil {
		return &[]models.Bed{}, err
	}

	var beds []models.Bed

	for rows.Next() {
		var bed models.Bed
		if err = rows.Scan(&bed.Id, &bed.Bed_Number, &bed.Price, &bed.Status, &bed.Room_id); err != nil {
			log.Fatal(err)
		}
		beds = append(beds, bed)
	}

	return &beds, nil
}