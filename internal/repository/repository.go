package repository

import "github.com/kylelaverty/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
