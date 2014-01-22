package evergreen

import (
	"errors"
	"net/http"
)

type ETag struct{}

func (self *ETag) Get(req *http.Request) (result string, err error) {
	result = req.Header.Get("If-None-Match")

	if result == "" {
		err = errors.New("Identity not found in ETag")
	}

	return
}
