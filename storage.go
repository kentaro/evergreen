package evergreen

import (
	"net/http"
)

type Storage interface {
	Get(req *http.Request) (result string, err error)
	Set(writer http.ResponseWriter, value string)
}
