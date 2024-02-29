package models

import "time"

// Reservation holds reservation data.
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the model for the user table.
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is the model for the rooms table.
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is the model for the restrictions table.
type Restrictions struct {
	ID          int
	Restriction string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Reservations is the model for the reservations table.
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RoomRestrictions is the model for the room_restrictions table.
type RoomRestrictions struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Reservations
	RestrictionID int
	Restriction   Restrictions
}
