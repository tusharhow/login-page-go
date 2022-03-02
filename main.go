package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Welcome struct {
	Name string
}

func main() {
	welcome := Welcome{Name: "Your Name"}

	templates := template.Must(template.ParseFiles("template/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening.. wait", string(welcome.Name))
	fmt.Println(http.ListenAndServe(":8080", nil))
}
