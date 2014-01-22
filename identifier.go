package evergreen

import (
	"net/http"
)

type Identifier interface {
	Get(req *http.Request)
	Set(writer http.ResponseWriter, value string)
}
