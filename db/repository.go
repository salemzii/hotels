package db

type Repository interface {
	Migrate() error
	Create()
	All()
	Delete(id int64) error
}
