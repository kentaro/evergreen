package evergreen

import (
	"net/http"
	"strings"
)

func newRequest() (req *http.Request) {
	req, _ = http.NewRequest("GET", "/test", strings.NewReader("test"))
	return
}
