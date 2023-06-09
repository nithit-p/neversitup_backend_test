package postgresql

import (
	"database/sql"
	"fmt"

	"user_management/domain"
	"user_management/repository"
)

// Verify Interface
var _ repository.UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewUserRepository(db *sql.DB, tableName string) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (repo *UserRepositoryImpl) ListUser() ([]domain.User, error) {
	var users []domain.User

	queryStr := fmt.Sprintf("SELECT id, username, email, first_name, last_name FROM %s ", repo.tableName)
	rows, err := repo.db.Query(queryStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepositoryImpl) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	// Execute the SQL statement
	queryStr := fmt.Sprintf("SELECT id, username, email, first_name, last_name FROM %s WHERE username=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, username).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUserByID(userId int) (domain.User, error) {
	var user domain.User
	// Execute the SQL statement
	queryStr := fmt.Sprintf("SELECT id, username, email, first_name, last_name FROM %s WHERE id=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, userId).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) CreateUser(username, email, firstName, lastName string) error {
	// Execute the SQL statement
	queryStr := fmt.Sprintf("INSERT INTO %s (username, email, first_name, last_name) VALUES ($1, $2, $3, $4)", repo.tableName)
	_, err := repo.db.Exec(queryStr, username, email, firstName, lastName)
	if err != nil {
		return err
	}

	return nil
}
