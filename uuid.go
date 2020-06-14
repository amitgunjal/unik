package uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	Hyphen         = "-"
	EmptyString    = ""
	ValNegativeOne = -1
)

func NewUUID() string {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err).String()
	// remove all dashes
	// -1 means, all occurrences
	u2 := strings.Replace(u1, Hyphen, EmptyString, ValNegativeOne)
	return u2
}
