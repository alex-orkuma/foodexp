package mysql

import (
	"database/sql"
	"errors"

	"github.com/alex-orkuma/foodexp/pkg/models"
)

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) Insert(food_id, food_name, shelf_life string) (int, error) {

	stm := `INSERT INTO products (food_id, food_name, shelf_life)
	VALUES ( ?, ?, ?)`

	res, err := m.DB.Exec(stm, food_id, food_name, shelf_life)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *ProductModel) Get(id int) (*models.Products, error) {

	stm := `SELECT id, food_id, food_name, shelf_life FROM products
	WHERE id = ?`

	row := m.DB.QueryRow(stm, id)

	s := &models.Products{}

	err := row.Scan(&s.ID, &s.FoodID, &s.FoodName, &s.ShelfLife)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *ProductModel) Latest() ([]*models.Products, error) {

	stm := `SELECT id, food_id, food_name, shelf_life FROM products
	ORDER BY id ASC LIMIT 20`

	rows, err := m.DB.Query(stm)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*models.Products{}

	for rows.Next() {
		p := &models.Products{}

		err := rows.Scan(&p.ID, &p.FoodID, &p.FoodName, &p.ShelfLife)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
