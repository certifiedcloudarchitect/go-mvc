package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("{Id: 1024, firstName: \"Git\", lastName: \"Hub\", email: \"github@gmail.com\"}"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
