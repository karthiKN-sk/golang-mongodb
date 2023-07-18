package main

import (
	"fmt"
	"golang-mongodb/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("users/:id", uc.GetUser)
	r.POST("users", uc.CreateUser)
	r.DELETE("users/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	// var uri string = "mongodb+srv://karthikeyanskalyan:KarthiSuvi@karthisuvilearning.isqslg6.mongodb.net/"
	// var uri string = "mongodb://localhost:27017  ==mongodb://karthikeyanskalyan:KarthiSuvi@127.0.0.1:27017/golang-mongodb"

	s, err := mgo.Dial("mongodb://localhost:27017/")

	if err != nil {
		fmt.Println("DB Not Connected", err)
		panic(err)
	}
	fmt.Println("DB Connected")
	return s
}
