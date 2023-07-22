package main

import (
	//"fmt"
	"html/template"
	"main/handlers"
	"main/utils"
	"net/http"
)

func init() {
	utils.Templ = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.ListenAndServe(":6600", nil)
}
