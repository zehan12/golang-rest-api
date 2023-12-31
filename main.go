package main

import (
	"fmt"
	"golang-rest-api/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mangodb://localhost:27107")
	if err == nil {
		fmt.Println(err)
		panic(err)
	}
	if s != nil {
		fmt.Println(s)
	}
	return s
}
