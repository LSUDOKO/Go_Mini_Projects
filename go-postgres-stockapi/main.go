package main

import (
	"fmt"
	"go-postgres-stockapi/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Satrting server on port 8081....")
	log.Fatal(http.ListenAndServe(":8081", r))
}
