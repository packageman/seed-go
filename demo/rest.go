package demo

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/ant0ine/go-json-rest/rest"
)

func MakeRest() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))

	log.Infof("Server started, listen on 8080 port")
	http.ListenAndServe(":8080", api.MakeHandler())
}
