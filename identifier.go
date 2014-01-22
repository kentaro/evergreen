package evergreen

import (
	"net/http"
)

type Identifier interface {
	Get(req *http.Request)
	Set(res http.ResponseWriter, value string)
}
