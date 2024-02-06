package item

import "github.com/google/uuid"

type Id uuid.UUID

func (e Id) UUID() uuid.UUID {
	return uuid.UUID(e)
}