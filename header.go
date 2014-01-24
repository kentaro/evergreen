package evergreen

import (
	"net/http"
)

type Header struct{}

func (self *Header) Set(writer http.ResponseWriter, value string) {
	writer.Header().Set("X-Evergreen-Id", value)
}
