package evergreen

import (
	"errors"
	"net/http"
)

type CookieDetector struct{}

func (self *CookieDetector) Detect(key string, req *http.Request) (result string, err error) {
	for _, cookie := range req.Cookies() {
		if cookie.Name == key {
			result = cookie.Value
			break
		}
	}

	if result == "" {
		err = errors.New("Identity not found in cookie")
	}

	return
}
