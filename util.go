package evergreen

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

func newUuid() (result string, err error) {
	u4, err := uuid.NewV4()

	if err != nil {
		log.Printf("UUID generation error: %v", err)
	}

	result = u4.String()
	return
}
