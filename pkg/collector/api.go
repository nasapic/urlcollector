/*
Package collector defines a generic interface to collect pictures from external
APIs.
For the sake of simplicity it is defined now within the service libs
but eventually it could be moved to a custom module in order to facilitateits
reuse.
*/
package collector

import (
	"time"
)

type (
	API interface {
		GetBetweenDates(from, to time.Time) (r Result, err error)
	}

	Result interface {
		GetList() []string
	}
)
