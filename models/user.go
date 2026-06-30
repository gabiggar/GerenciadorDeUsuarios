package models

import (
	dto "GerenciadorDeUsuarios/dto"
	"errors"

	"github.com/google/uuid"
)

type id uuid.UUID

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	biography string
}

type application struct {
	data map[id]user
}

var ErrNotFound = errors.New("usuário não encontrado")

var db = application{
	data: map[id]user{
		id(uuid.New()): {
			FirstName: "Gabriela",
			LastName:  "Garcia",
			biography: "Backend Developer",
		},
		id(uuid.New()): {
			FirstName: "Augusto",
			LastName:  "Silva",
			biography: "Frontend Developer",
		},
	},
}

func FindAll() ([]dto.UserResponse, error) {
	users := make([]dto.UserResponse, 0, len(db.data))

	for id, user := range db.data {
		users = append(users, dto.UserResponse{
			ID:        uuid.UUID(id).String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Biography: user.biography,
		})
	}

	return users, nil
}

func FindById(userId uuid.UUID) (*dto.UserResponse, error) {
	user, ok := db.data[id(userId)]

	if !ok {
		return &dto.UserResponse{}, ErrNotFound
	}

	uResponse := dto.UserResponse{
		ID:        userId.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.biography,
	}

	return &uResponse, nil
}

func Insert(newUser dto.CreateUserDTO) (dto.UserResponse, error) {

	newID := uuid.New()

	u := user{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		biography: newUser.Biography,
	}

	db.data[id(newID)] = u

	return dto.UserResponse{
		ID:        newID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Biography: u.biography,
	}, nil

}

func Update(userId uuid.UUID, userUpdates dto.UpdateUserDTO) (*dto.UserResponse, error) {
	storedUser, ok := db.data[id(userId)]
	if !ok {
		return nil, ErrNotFound
	}

	storedUser.FirstName = userUpdates.FirstName
	storedUser.LastName = userUpdates.LastName
	storedUser.biography = userUpdates.Biography

	db.data[id(userId)] = storedUser

	userResponse := &dto.UserResponse{
		ID:        userId.String(),
		FirstName: storedUser.FirstName,
		LastName:  storedUser.LastName,
		Biography: storedUser.biography,
	}

	return userResponse, nil
}

func Delete(userId uuid.UUID) (*dto.UserResponse, error) {
	parsedUserId := id(userId)
	storedUser, ok := db.data[parsedUserId]
	if !ok {
		return nil, ErrNotFound
	}

	userResponse := &dto.UserResponse{
		ID:        userId.String(),
		FirstName: storedUser.FirstName,
		LastName:  storedUser.LastName,
		Biography: storedUser.biography,
	}

	delete(db.data, parsedUserId)

	return userResponse, nil
}
