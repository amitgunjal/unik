package unik

import (
	"strings"

	satoriUUID "github.com/satori/go.uuid"
)

const (
	Hyphen         = "-"
	EmptyString    = ""
	ValNegativeOne = -1
)

func NewUUID() string {
	id, _ := satoriUUID.NewV4()
	var err error
	u1 := satoriUUID.Must(id, err).String()
	// remove all dashes
	// -1 means, all occurrences
	u2 := strings.Replace(u1, Hyphen, EmptyString, ValNegativeOne)
	return u2
}
