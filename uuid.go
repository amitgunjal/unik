package unik

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	Hyphen         = "-"
	EmptyString    = ""
	ValNegativeOne = -1
)

func NewUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("unable to get uuid.")
	}
	u1 := uuid.Must(id, err).String()
	// remove all dashes
	// -1 means, all occurrences
	u2 := strings.Replace(u1, Hyphen, EmptyString, ValNegativeOne)
	return u2
}
