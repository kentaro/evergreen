package evergreen

import (
	"log"
	"net/http"
)

type Server struct {}

func (self *Server) Run(options map[string]string) {
	identifiers := []Identifier{&Cookie{key:options["cookie_key"]}, &ETag{}}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		var value string

		for _, identity := range identifiers {
			result, err := identity.Get(req)

			if err == nil {
				value = result
				break
			}
		}

		if value == "" {
			uuid, err := newUuid()

			if err != nil {
				return
			}

			value = uuid
		}

		for _, identity := range identifiers {
			identity.Set(writer, value)
		}
	})

	log.Fatal(http.ListenAndServe(":" + options["port"], nil))
}
