package service

import (
	"hotels/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupApi() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mybuks.netlify.app", "http://localhost:5174", "http://localhost:8891"}

	//config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"POST", "PUT", "GET", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}

	router.Use(cors.New(config))

	router.POST("/api/create/hotel", Createhotel)
	router.GET("/api/get/hotels", AllHotels)
	router.GET("/api/get/hotel/:id", GetHotel)

	router.GET("/api/get/floors", AllFloors)
	router.GET("/api/get/floor/:id", FloorById)
	router.GET("/api/get/floors/:hotelid", FloorByHotel)

	router.POST("/api/create/room", CreateRoom)
	router.GET("/api/get/room/:id", RoomById)
	router.GET("/api/get/rooms/:floorid", RoomByFloor)

	router.POST("/api/create/bed", CreateBed)
	router.GET("/api/get/beds", AllBeds)
	router.GET("/api/get/bed/:id", BedById)
	router.GET("/api/get/beds/:roomid", BedByRoom)
	router.GET("/api/get/beds/available/:roomid", AvailableBeds)
	router.GET("/api/get/beds/occupied/:roomid", OccupiedBeds)
	router.GET("/api/get/beds/checkout/:roomid", CheckOutBeds)
	router.PUT("/api/update/bed/:id")

	router.POST("/api/create/reservation", CustomerBooking)

	return router
}

func hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Hello salemzii")
	}
}

/* begin hotels */

// fetch all hotels
func AllHotels(c *gin.Context) {
	hotels, err := HotelRepository.FetchAllHotel()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hotels": hotels,
	})
}

// fetch single hotel by :id
func GetHotel(c *gin.Context) {
	hotelId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
	}

	id, _ := strconv.Atoi(hotelId)
	hotel, err := HotelRepository.FetchHotel(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"hotel": hotel,
	})
}

// create new hotel
func Createhotel(c *gin.Context) {
	var createhotel models.CreateHotel
	var hotel models.Hotel
	var createhotelResp models.CreateHotel

	if err := c.ShouldBindJSON(&createhotel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	hotel.Name = createhotel.Name
	hotel.Address = createhotel.Address
	hotel.SetPrefix()

	createdHotel, err := HotelRepository.CreateHotel(&hotel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if createdHotel.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "object id cannot be 0",
		})
		return
	} else {
		createhotelResp.Id = createdHotel.Id
		createhotelResp.Name = createdHotel.Name
		createhotelResp.Address = createdHotel.Address

		for _, v := range createhotel.Floors {
			v.Hotel_id = createdHotel.Id
			createdFloor, err := HotelRepository.CreateFloor(&v)
			if err != nil {
				log.Println(err)
			}
			createhotelResp.Floors = append(createhotelResp.Floors, *createdFloor)
		}

		c.JSON(http.StatusCreated, gin.H{
			"hotel": createhotelResp,
		})
	}
}

/* end hotels */

/* begin floors */

func AllFloors(c *gin.Context) {
	floors, err := HotelRepository.FetchAllFloor()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"floors": floors,
	})
}

func FloorById(c *gin.Context) {
	floorId, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(floorId)
	floor, err := HotelRepository.FetchFloor(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"floor": floor,
	})
}

func FloorByHotel(c *gin.Context) {
	hotelId, ok := c.Params.Get("hotelid")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(hotelId)
	floor, err := HotelRepository.FetchFloorsByHotel(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"floor": floor,
	})
}

/* end floors */

/* begin rooms */

// create room
func CreateRoom(c *gin.Context) {
	var room models.Room

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdRoom, err := HotelRepository.CreateRoom(&room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"room": createdRoom,
	})
}

// get all rooms
func AllRooms(c *gin.Context) {
	rooms, err := HotelRepository.GetAllRoom()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

// get room by :id
func RoomById(c *gin.Context) {
	roomId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(roomId)

	room, err := HotelRepository.GetRoom(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"room": room,
	})
}

// get room by floorId
func RoomByFloor(c *gin.Context) {
	floorId, ok := c.Params.Get("floorid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(floorId)
	rooms, err := HotelRepository.FetchRoomsByFloor(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

/* end rooms */

/* begin beds */

func CreateBed(c *gin.Context) {
	var bed models.Bed
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdBed, err := HotelRepository.CreateBed(&bed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"bed": createdBed,
	})
}

func AllBeds(c *gin.Context) {
	beds, err := HotelRepository.FetchAllBed()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beds": beds,
	})
}

func BedById(c *gin.Context) {
	bedId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(bedId)

	bed, err := HotelRepository.FetchBed(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"bed": bed,
	})
}

func BedByRoom(c *gin.Context) {
	roomId, ok := c.Params.Get("roomid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(roomId)
	beds, err := HotelRepository.FetchBedsByRoom(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beds": beds,
	})
}

func AvailableBeds(c *gin.Context) {
	roomId, ok := c.Params.Get("roomid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(roomId)
	beds, err := HotelRepository.FetchAvailableBeds(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beds": beds,
	})
}

func OccupiedBeds(c *gin.Context) {
	roomId, ok := c.Params.Get("roomid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(roomId)
	beds, err := HotelRepository.FetchOccupiedBeds(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beds": beds,
	})
}

func CheckOutBeds(c *gin.Context) {
	roomId, ok := c.Params.Get("roomid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not found",
		})
		return
	}

	id, _ := strconv.Atoi(roomId)
	beds, err := HotelRepository.FetchCheckOutBeds(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beds": beds,
	})
}

/* end beds */

/* begin customers */

func CustomerBooking(c *gin.Context) {
	var Booking models.CreateCustomer
	var customer models.Customer
	var BookingResp models.CreateCustomer

	if err := c.ShouldBindJSON(&Booking); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.Name = Booking.Name
	customer.Agent_name = Booking.Agent_name

	createdCustomer, err := HotelRepository.CreateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if createdCustomer.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "customer id cannot be 0",
		})
		return
	} else {
		BookingResp.Id = createdCustomer.Id
		BookingResp.Name = createdCustomer.Name
		BookingResp.Agent_name = createdCustomer.Agent_name

		for _, v := range Booking.Reservations {
			v.Customer_id = createdCustomer.Id

			createdRSV, err := HotelRepository.CreateReservation(&v)
			if err != nil {
				log.Println(err)
			}
			BookingResp.Reservations = append(BookingResp.Reservations, *createdRSV)
		}

		c.JSON(http.StatusCreated, gin.H{
			"resp": BookingResp,
		})
	}

}

/* end customers */
