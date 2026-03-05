package middleware

import (
	"database/sql"
	"datablse/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type response struct {
	ID      string `json:"id,omitempty`
	Message string `json:"message,omitempty`
}

func CreateConnection() *sql.DB {
	err := gotdotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucessfull connected to postgress")
	return db
}

func GetStock() {

}

func GetAllStock() {

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.stock
	json.NewDecoder(r.Body).Decode(&stock)

}

func UpdateStock() {

}
func DeleteStock() {

}
