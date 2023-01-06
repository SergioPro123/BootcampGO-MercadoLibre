package interfaces

import (
	"proyectoapisupermercado/models"
)

type IDatabaseProduct interface {
	GetProducts() ([]models.Product, error)
	GetLastId() (int, error)
}
