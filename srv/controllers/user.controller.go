package controller

import (
	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"text/template"
)

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		ageStr := r.FormValue("age") // get Age as string

		// Convert the string ageStr to int
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age value", http.StatusBadRequest)
			return
		}

		user := models.User{Name: name, Email: email, Age: age} // use age how int
		err = models.CreateUser(db, &user)
		if err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/users", http.StatusFound)
	} else {
		tmpl := template.Must(template.ParseFiles("views/create_user.html"))
		tmpl.Execute(w, nil) // Render template to create user
	}
}
