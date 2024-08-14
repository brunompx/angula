package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/brunompx/go-delicatessen/handlers"
	"github.com/brunompx/go-delicatessen/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	cfg := mysql.Config{
		User:                 storage.Envs.DBUser,
		Passwd:               storage.Envs.DBPassword,
		Addr:                 storage.Envs.DBAddress,
		DBName:               storage.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := storage.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)

	initStorage(db)

	router := mux.NewRouter()

	handler := handlers.New(store)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")

	//router.HandleFunc("/cars", handler.HandleListCars).Methods("GET")
	//router.HandleFunc("/cars", handler.HandleAddCar).Methods("POST")
	//router.HandleFunc("/cars/{id}", handler.HandleDeleteCar).Methods("DELETE")
	//router.HandleFunc("/cars/search", handler.HandleSearchCar).Methods("GET")

	router.HandleFunc("/products", handler.HandleListProducts).Methods("GET")
	router.HandleFunc("/products", handler.HandleAddProduct).Methods("POST")
	router.HandleFunc("/products/search", handler.HandleSearchProduct).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")
}
