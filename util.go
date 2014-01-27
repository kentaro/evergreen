package evergreen

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"strings"
)

func newUuid() (result string, err error) {
	u4, err := uuid.NewV4()

	if err != nil {
		log.Printf("UUID generation error: %v", err)
	}

	result = u4.String()
	return
}

func newRequest() (req *http.Request) {
	req, _ = http.NewRequest("GET", "/test", strings.NewReader("test"))
	return
}
