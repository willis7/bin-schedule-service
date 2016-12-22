package services

import (
	"github.com/willis7/bin-schedule-service/models"
)

// Client provides an interface to the schedule package
type Client interface {
	Get(string)([]models.Schedule, error)
}
