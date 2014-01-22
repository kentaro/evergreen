package evergreen

import (
	"errors"
	"net/http"
)

type Cookie struct {
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

func (self *Cookie) Set(writer http.ResponseWriter, value string) {
	cookie := http.Cookie{
		Name:       self.key,
		Value:      value,
		RawExpires: "Tue, 31 Dec 2030 23:30:45 GMT",
		MaxAge:     630720000,
	}
	writer.Header().Set("Set-Cookie", cookie.String())

	return
}
