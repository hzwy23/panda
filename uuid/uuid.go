package uuid

import (
	"github.com/satori/go.uuid"
)

func Random() string {
	return uuid.NewV4().String()
}

func UUID() string {
	uid1 := uuid.NewV1().String()
	uid2 := uuid.NewV4()
	return uuid.NewV5(uid2, uid1).String()
}
