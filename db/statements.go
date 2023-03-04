package db

var (
	CreateHotel = `
		CREATE TABLE IF NOT EXISTS hotels (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			address TEXT NOT NULL,
			prefix TEXT NOT NULL
		) 
	`
	CreateFloor = `
		CREATE TABLE IF NOT EXISTS floors (
			id SERIAL PRIMARY KEY, 
			floor_no INT NOT NULL,
			hotel_id integer REFERENCES hotels(id)
		)
	`
	CreateRoom = `
		CREATE TABLE IF NOT EXISTS rooms (
			id SERIAL PRIMARY KEY,
			room_no INT NOT NULL,
			floor_id integer REFERENCES floors(id)
		)
	`
	CreateBed = `
		CREATE TABLE IF NOT EXISTS beds (
			id SERIAL PRIMARY KEY,
			bed_no VARCHAR(255) NOT NULL,
			price FLOAT NOT NULL,
			status VARCHAR(25) NOT NULL,
			room_id integer REFERENCES rooms(id)
		)			 
	`
	CreateCustomer = `
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(75),
			agent_name VARCHAR(75)
		)
	`
	CreateCustomerReservation = `
		CREATE TABLE IF NOT EXISTS customer_reservations (
			id SERIAL PRIMARY KEY,
			from_date TIMESTAMP,
			to_date TIMESTAMP,
			bed_id integer REFERENCES beds(id) UNIQUE,
			customer_id integer REFERENCES customers(id) UNIQUE
		)
	`

	InsertHotel    = `INSERT INTO hotels (name, address, prefix) VALUES($1, $2, $3) RETURNING id, name, address, prefix`
	InsertFloor    = `INSERT INTO floors (floor_no, hotel_id) VALUES($1, $2) RETURNING id, floor_no, hotel_id`
	InsertRoom     = `INSERT INTO rooms (room_no, floor_id) VALUES($1, $2) RETURNING id, room_no, floor_id`
	InsertBed      = `INSERT INTO beds (bed_no, price, status, room_id) VALUES($1, $2, $3, $4) RETURNING id, bed_no, price, status, room_id`
	InsertCustomer = `INSERT INTO customers (name, agent_name) 
							VALUES($1, $2) 
							RETURNING id, name, agent_name`
	InsertCustomerReservation = `INSERT INTO customer_reservations (from_date, to_date, bed_id, customer_id) 
									VALUES($1, $2, $3, $4) 
									RETURNING id, from_date, to_date, bed_id, customer_id`

	InsertFloorMany               = ``
	InsertCustomerReservationMany = ``

	FetchAllHotels       = `SELECT * FROM hotels`
	FetchAllFloors       = `SELECT * FROM floors`
	FetchAllRooms        = `SELECT * FROM rooms`
	FetchAllBeds         = `SELECT * FROM beds`
	FetchAllCustomers    = `SELECT * FROM customers`
	FetchAllReservations = `SELECT * FROM customer_reservations`

	FetchHotelById       = `SELECT * FROM hotels WHERE id=($1)`
	FetchFloorById       = `SELECT * FROM floors WHERE id=($1)`
	FetchRoomById        = `SELECT * FROM rooms WHERE id=($1)`
	FetchBedById         = `SELECT * FROM beds WHERE id=($1)`
	FetchCustomerById    = `SELECT * FROM customers WHERE id=($1)`
	FetchReservationById = `SELECT * FROM customer_reservations WHERE id=($1)`

	FetchFloorByHotelId          = `SELECT * FROM floors WHERE hotel_id=($1)`
	FetchRoomByFloorId           = `SELECT * FROM rooms WHERE floor_id=($1)`
	FetchBedByRoomId             = `SELECT * FROM beds WHERE room_id=($1)`
	FetchReservationByCustomerId = `SELECT * FROM customer_reservations WHERE customer_id=($1)`
	FetchCustomerReservations = `SELECT * FROM customers JOIN customer_reservations ON ($1)=customer_id;`

	FetchAvailableBedsByRoomId = `SELECT * FROM beds WHERE room_id=($1) AND status='cleaned'`
	FetchOccupiedBedsByRoomId  = `SELECT * FROM beds WHERE room_id=($1) AND status='occupied'`
	FetchCheckoutBedsByRoomId  = `SELECT * FROM beds WHERE room_id=($1) AND status='checkout'`


	UpdateBedStatus = `
		UPDATE beds
		SET status=($2)
		WHERE room_id=($1);
	`
)
