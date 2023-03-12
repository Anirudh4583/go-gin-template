package database

import (
	"fmt"
)

func Setup(){
    fmt.Println("working setup()")
	db:= ConnectToDb()
    defer db.Close()

    err := db.DB().Ping()
    if err != nil {
        fmt.Printf("Error pinging the database: %v\n", err)
        return
    }

    fmt.Println("Connected to the database!")    
}