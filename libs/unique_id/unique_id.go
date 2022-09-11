package unique_id

import "github.com/google/uuid"

func New() string {
	return uuid.NewString()
}
