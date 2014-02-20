package evergreen

import (
	"log"
	"net/http"
)

type Server struct {}

func (self *Server) Run(options map[string]string) {
	storages := []Storage{&Cookie{Key:options["cookie_key"]}, &ETag{}}
	header := &Header{}
	emptyGif := &EmptyGif{}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		var value string

		for _, storage := range storages {
			result, err := storage.Get(req)

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

		for _, storage := range storages {
			storage.Set(writer, value)
		}

		header.Set(writer, value)
		emptyGif.Set(writer)
	})

	log.Fatal(http.ListenAndServe(":" + options["port"], nil))
}
