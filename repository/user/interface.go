package user

import "Alterra/batch5/docker-be5/entities"

type User interface {
	Get() ([]entities.User, error)
}
