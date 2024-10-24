package controller

import (
	"net/http"

	"gorm.io/gorm"
)

func HelloWorld(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Set the response content type to HTML
		w.Header().Set("Content-Type", "text/html")

		html := `
            <!DOCTYPE html>
            <html lang="pt-BR">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Hello World</title>
            </head>
            <body>
                <h1>Hello World</h1>
                <p>Welcome to Go Application</p>
            </body>
            </html>
        `
		w.Write([]byte(html))
		return
	}

	//For other methods, you can return a method not allowed.
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed."))
}
