package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprint(w, "Welcome, "+r.Form.Get("name"))
}
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	fmt.Fprintf(w, "you are get user %s", userId)
}
func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	fmt.Fprintf(w, "you are update user %s", userId)
}
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	fmt.Fprintf(w, "you are delete user %s", userId)
}
func Store(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "you are store user.")
}
func main() {
	router := httprouter.New()
	router.GET("/users/", Index)
	router.POST("/users/", Store)
	router.GET("/users/:user_id", Show)
	router.PUT("/users/:user_id", Update)
	router.DELETE("/users/:user_id", Delete)

	log.Fatal(http.ListenAndServe(":8088", router))
}
