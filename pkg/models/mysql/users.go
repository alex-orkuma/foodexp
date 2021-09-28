package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/alex-orkuma/foodexp/pkg/models"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(first_name, last_name, email, password, role string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users_admin (first_name, last_name, email, password, role, statuss)
	VALUES(?, ?, ?, ?, ?, "active")`

	_, err = m.DB.Exec(stmt, first_name, last_name, email, string(hashedPassword), role)
	if err != nil {

		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err

	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
