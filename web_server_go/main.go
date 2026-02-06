package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method NOt allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request sucessfull\n")
	name := r.FormValue("name")
	Roll_NO := r.FormValue("Roll_NO")
	Branch := r.FormValue("Branch")
	Room_No := r.FormValue("Room_No")
	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Rollno %s\n", Roll_NO)
	fmt.Fprintf(w, "Branch%s\n", Branch)
	fmt.Fprintf(w, "Room No %s\n", Room_No)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server start at port 8081")
	fmt.Println()
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
