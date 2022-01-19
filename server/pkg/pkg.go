package pkg

import "github.com/google/uuid"

func Hoge() (uuid.UUID, error) {
	return uuid.NewRandom()
}
