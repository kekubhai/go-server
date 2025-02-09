package main

import (
    "log"
    "net/http"
      "github.com/akhil/mongo-golang/controllers"
    "github.com/julienschmidt/httprouter"
    
    "gopkg.in/mgo.v2"
)

func main() {
    session, err := mgo.Dial("mongodb://localhost:27017")
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }
    defer session.Close()

    uc := controllers.NewUserController(session)
    r := httprouter.New()
    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.DeleteUser)

    log.Println("Server is running on port 9000")
    if err := http.ListenAndServe(":9000", r); err != nil {
        log.Fatal("Server failed:", err)
    }
}
