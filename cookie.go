package evergreen

import (
	"errors"
	"net/http"
	"time"
)

type Cookie struct {
	Key string
}

func (self *Cookie) Get(req *http.Request) (result string, err error) {
	for _, cookie := range req.Cookies() {
		if cookie.Name == self.Key {
			result = cookie.Value
			break
		}
	}

	if result == "" {
		err = errors.New("Identity not found in cookie")
	}

	return
}

var expires, _ = time.Parse("Mon, 2 Jan 2006 15:04:05 MST", "Tue, 31 Dec 2030 23:30:45 GMT")

func (self *Cookie) Set(writer http.ResponseWriter, value string) {
	cookie := http.Cookie{
		Name:    self.Key,
		Value:   value,
		Expires: expires,
		MaxAge:  630720000,
	}
	writer.Header().Set("Set-Cookie", cookie.String())

	return
}
