package evergreen

import (
	"errors"
	"net/http"
)

type Cookie struct{
	key string
}

func (self *Cookie) Get(req *http.Request) (result string, err error) {
	for _, cookie := range req.Cookies() {
		if cookie.Name == self.key {
			result = cookie.Value
			break
		}
	}

	if result == "" {
		err = errors.New("Identity not found in cookie")
	}

	return
}
