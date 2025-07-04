package simpleserver

import (
	"fmt"
	"net/http"
)

func MyServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Go!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./functions/simpleserver/home.html")
	})

	http.ListenAndServe(":3000", nil)
}
