package usercontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fahrigunadi/go-crud/entities"
	"github.com/fahrigunadi/go-crud/libraries"
	"github.com/fahrigunadi/go-crud/models"
)

var validation = libraries.NewValidation()
var userModel = models.NewUserModel()

func Index(res http.ResponseWriter, req *http.Request) {

	users, _ := userModel.FindAll()

	data := map[string]interface{}{
		"users": users,
	}

	tmp, err := template.ParseFiles("views/users/index.html")

	if err != nil {
		panic(err)
	}

	tmp.Execute(res, data)
}

func Add(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {

		tmp, err := template.ParseFiles("views/users/add.html")

		if err != nil {
			panic(err)
		}

		tmp.Execute(res, nil)

	} else if req.Method == http.MethodPost {

		req.ParseForm()

		var user entities.User
		user.Name = req.Form.Get("name")
		user.Username = req.Form.Get("username")
		user.Phone = req.Form.Get("phone")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["user"] = user
			data["validation"] = vErrors
		} else {
			data["message"] = "User Created Successfully!"
			userModel.Create(user)
		}

		tmp, _ := template.ParseFiles("views/users/add.html")
		tmp.Execute(res, data)
	}
}

func Edit(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {

		queryString := req.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var user entities.User
		userModel.Find(id, &user)

		data := map[string]interface{}{
			"user": user,
		}

		tmp, err := template.ParseFiles("views/users/edit.html")

		if err != nil {
			panic(err)
		}

		tmp.Execute(res, data)

	} else if req.Method == http.MethodPost {

		req.ParseForm()

		var user entities.User
		user.Id, _ = strconv.ParseInt(req.Form.Get("id"), 10, 64)
		user.Name = req.Form.Get("name")
		user.Username = req.Form.Get("username")
		user.Phone = req.Form.Get("phone")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["user"] = user
			data["validation"] = vErrors
		} else {
			data["message"] = "User Updated Successfully!"
			userModel.Update(user)
		}

		tmp, _ := template.ParseFiles("views/users/edit.html")
		tmp.Execute(res, data)
	}
}

func Delete(res http.ResponseWriter, req *http.Request) {
	queryString := req.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(res, req, "/users", http.StatusSeeOther)
}
