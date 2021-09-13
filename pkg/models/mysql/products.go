package mysql

import (
	"database/sql"

	"github.com/alex-orkuma/foodexp/pkg/models"
)

type ProductModel struct {
	DB sql.DB
}

func (m *ProductModel) Insert(foodname, shelflife, foodid string) (int, error) {
	return 0, nil
}

func (m *ProductModel) Get(id int) (*models.Products, error) {
	return nil, nil
}

func (m *ProductModel) Latest() ([]*models.Products, error) {
	return nil, nil
}
