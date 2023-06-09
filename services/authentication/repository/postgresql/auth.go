package postgresql

import (
	"database/sql"
	"errors"
	"fmt"

	"authentication/domain"
	"authentication/repository"

	"golang.org/x/crypto/bcrypt"
)

// Verify Interface
var _ repository.AuthRepository = (*AuthRepositoryImpl)(nil)

type AuthRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewAuthRepository(db *sql.DB, tableName string) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (repo *AuthRepositoryImpl) Login(username string, password string) error {
	var hashedPassword string

	// Execute the SQL statement
	queryStr := fmt.Sprintf("SELECT password_hash FROM %s WHERE username=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	// Compare the password and hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

func (repo *AuthRepositoryImpl) GetAuthByUsername(username string) (domain.Auth, error) {
	var auth domain.Auth
	// Execute the SQL statement
	queryStr := fmt.Sprintf("SELECT auth_id, user_id, username, email FROM %s WHERE username=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, username).Scan(&auth.AuthId, &auth.UserId, &auth.Username, &auth.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return auth, nil
		}
		return auth, err
	}

	return auth, nil
}

func (repo *AuthRepositoryImpl) GetAuthByID(id int) (domain.Auth, error) {
	var auth domain.Auth
	// Execute the SQL statement
	queryStr := fmt.Sprintf("SELECT auth_id, user_id, username, email FROM %s WHERE id=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, id).Scan(&auth.AuthId, &auth.UserId, &auth.Username, &auth.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return auth, nil
		}
		return auth, err
	}

	return auth, nil
}

func (repo *AuthRepositoryImpl) CreateAuth(userId int, username, email, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Execute the SQL statement
	queryStr := fmt.Sprintf("INSERT INTO %s (user_id, username, email, password_hash) VALUES ($1, $2, $3, $4)", repo.tableName)
	_, err = repo.db.Exec(queryStr, userId, username, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
