package main
import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/LSUDOKO/Go_Mini_Projects/tree/main/Books_Managment_Mysql/pkg/routes"


)
func main(){
	r:=mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r ))
}