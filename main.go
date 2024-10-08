package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brunompx/angula/handlers"
	"github.com/brunompx/angula/storage"
	mysqld "github.com/go-sql-driver/mysql"
	mysqlg "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	/*cfg := mysqld.Config{
		User:                 storage.Envs.DBUser,
		Passwd:               storage.Envs.DBPassword,
		Addr:                 storage.Envs.DBAddress,
		DBName:               storage.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}*/

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "postgres", "postgres", "5432")

	//db, err := NewSQLStorage(cfg)
	db, err := NewPostgresSQLStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)

	initStorage(db)

	handler := handlers.New(store)

	//Using gorilla/mux
	//router := mux.NewRouter()
	//router.HandleFunc("/", handler.HandleHome).Methods("GET")
	//router.HandleFunc("/products", handler.HandleListProducts).Methods("GET")
	//router.HandleFunc("/products", handler.HandleAddProduct).Methods("POST")
	//router.HandleFunc("/products/search", handler.HandleSearchProduct).Methods("GET")
	//router.HandleFunc("/products/{id}", handler.HandleDeleteProduct).Methods("DELETE")
	// serve files in public
	//router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//fmt.Printf("Listening on %v\n", "localhost:8080")
	//http.ListenAndServe(":8080", router)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handler.HandleHome)

	router.HandleFunc("GET /products", handler.HandleListProducts)
	router.HandleFunc("POST /products", handler.HandleAddProduct)
	router.HandleFunc("GET /products/search", handler.HandleSearchProduct)
	router.HandleFunc("DELETE /products/{productID}", handler.HandleDeleteProduct)

	router.HandleFunc("GET /orders", handler.HandleListOrders)
	router.HandleFunc("POST /orders", handler.HandleAddOrder)
	router.HandleFunc("PATCH /orders", handler.HandleUpdateOrder)
	router.HandleFunc("GET /orders/edit", handler.HandleEditOrder)
	router.HandleFunc("DELETE /orders/{orderID}", handler.HandleDeleteOrder)
	router.HandleFunc("POST /orders/orderitem/{productID}", handler.HandleAddOrderItem)
	router.HandleFunc("DELETE /orders/orderitem/{productID}", handler.HandleRemoveOrderItem)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func NewSQLStorage(cfg mysqld.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysqlg.Open(cfg.FormatDSN()), &gorm.Config{})
	return db, err
}

func NewPostgresSQLStorage(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func initStorage(db *gorm.DB) {
	genericDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	pingErr := genericDB.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")
}
