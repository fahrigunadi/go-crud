package main

import (
	"net/http"

	"github.com/fahrigunadi/go-crud/controllers/usercontroller"
)

func main() {
	http.HandleFunc("/users", usercontroller.Index)
	http.HandleFunc("/users/add", usercontroller.Add)
	http.HandleFunc("/users/edit", usercontroller.Edit)
	http.HandleFunc("/users/delete", usercontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
