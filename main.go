package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	db, err := gorm.Open("sqlite3", "restaurant.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	router := mux.NewRouter()
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/", home)

	fmt.Println("Hello World!")

	log.Fatal(http.ListenAndServe(":8080", router))

}
