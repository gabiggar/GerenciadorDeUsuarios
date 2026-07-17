package repository

import (
	"database/sql"
	"errors"

	"GerenciadorDeUsuarios/internal/models"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{db: db}
}

type Store interface {
	FindAll() ([]models.User, error)
	FindById(userID uuid.UUID) (*models.User, error)
	Insert(newUser models.CreateUserRequest) (models.User, error)
	Update(userID uuid.UUID, update models.UpdateUserRequest) (*models.User, error)
	Delete(userId uuid.UUID) (*models.User, error)
}

var ErrNotFound = errors.New("user not found")

func (s *store) FindAll() ([]models.User, error) {
	query := `SELECT id, first_name, last_name, biography FROM users`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var u models.User
		var idStr string

		err := rows.Scan(
			&idStr,
			&u.FirstName,
			&u.LastName,
			&u.Biography,
		)
		if err != nil {
			return nil, err
		}

		parsedID, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		u.ID = parsedID

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *store) FindById(userID uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, biography 
		FROM users
		WHERE id = ?
	`

	var u models.User
	var idStr string

	err := s.db.QueryRow(query, userID).Scan(
		&idStr,
		&u.FirstName,
		&u.LastName,
		&u.Biography,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	u.ID = parsedID

	return &u, nil
}

func (s *store) Insert(newUser models.CreateUserRequest) (models.User, error) {
	newID := uuid.New()

	u := models.User{
		ID:        newID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Biography: newUser.Biography,
	}

	query := `
		INSERT INTO users (id, first_name, last_name, biography)
		VALUES (?, ?, ?, ?)
	`

	if _, err := s.db.Exec(query, u.ID.String(), u.FirstName, u.LastName, u.Biography); err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (s *store) Update(userID uuid.UUID, userUpdate models.UpdateUserRequest) (*models.User, error) {

	user, err := s.FindById(userID)
	if err != nil {
		return nil, err
	}

	if userUpdate.FirstName != nil {
		user.FirstName = *userUpdate.FirstName
	}

	if userUpdate.LastName != nil {
		user.LastName = *userUpdate.LastName
	}

	if userUpdate.Biography != nil {
		user.Biography = *userUpdate.Biography
	}

	query := `
		UPDATE users
		SET first_name = ?, last_name = ?, biography = ?
		WHERE id = ?
	`

	_, err = s.db.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Biography,
		user.ID,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *store) Delete(userId uuid.UUID) (*models.User, error) {
	user, err := s.FindById(userId)
	if err != nil {
		return nil, err
	}

	query := `
		DELETE from users
		WHERE id = ?
	`

	_, err = s.db.Exec(query, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
