package u

import "github.com/google/uuid"

func GetUid() string {
	return uuid.NewString()
}
