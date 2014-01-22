package evergreen

import (
	"errors"
	"net/http"
)

type ETagDetector struct{}

func (self *ETagDetector) Detect(key string, req *http.Request) (result string, err error) {
	result = req.Header.Get("If-None-Match")

	if result == "" {
		err = errors.New("Identity not found in ETag")
	}

	return
}
