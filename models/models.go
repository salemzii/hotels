package models

import "time"

type Hotel struct {
	Id      int
	Name    string `json:"name"`
	Address string `json:"address"`
	Prefix  string
}

func (h *Hotel) SetPrefix() {
	prx := h.Name[:4]
	h.Prefix = prx
}

type CreateHotel struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Floors  []Floor `json:"floors"`
}

type Floor struct {
	Id           int
	Floor_Number int `json:"floor_no"`
	Hotel_id     int `json:"hotel_id"`
}

type Room struct {
	Id          int
	Room_Number int `json:"room_no"`
	Floor_id    int `json:"floor_id"`
}

type Bed struct {
	Id         int
	Bed_Number int     `json:"bed_no"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
	Room_id    int     `json:"room_id"`
}

type Customer struct {
	Id         int
	Name       string `json:"name"`
	Agent_name string `json:"agent_name"`
}

type Reservation struct {
	Id          int
	From_date   time.Time `json:"from_date"`
	To_date     time.Time `json:"to_date"`
	Bed_id      int       `json:"bed_id"`
	Customer_id int       `json:"customer_id"`
	Price       float64   `json:"price"`
}

type CreateCustomer struct {
	Id           int
	Name         string        `json:"name"`
	Agent_name   string        `json:"agent_name"`
	Reservations []Reservation `json:"reservations"`
}

type CustomerReservation struct {
	CustomerId    int
	Name          string `json:"name"`
	Agent_name    string `json:"agent_name"`
	ReservationId int
	From_date     time.Time `json:"from_date"`
	To_date       time.Time `json:"to_date"`
	Bed_id        int       `json:"bed_id"`
	Customer_id   int       `json:"customer_id"`
	Price         float64   `json:"price"`
}
