package models

import "github.com/google/uuid"

type Sandwich struct {
	id     uuid.UUID
	Name   string
	UserID uuid.UUID
	Desc   string
}
