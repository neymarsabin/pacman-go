package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// information on database credentials
const (
	host     = "localhost"
	port     = 5432
	user     = "prasvin"
	password = "gopacman123"
	dbname   = "postgres"
)

type Player struct {
	Id       int    `json:"id"`
	username string `json:"username"`
	score    int    `json:"score"`
}

func main() {
	psqlInfo := connectionWithPostres()
	fmt.Println(psqlInfo)
	http.HandleFunc("/", displayUserInfo)
	http.HandleFunc("/create_score", createUserScore)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func displayUserInfo(w http.ResponseWriter, r *http.Request) {
	p := Player{
		Id:       1,
		username: "neymarsabin",
		score:    20,
	}
	RespondWithJSON(w, http.StatusOK, p)
}

func createUserScore(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	decodeJson := json.NewDecoder(r.Body)
	fmt.Println(decodeJson)
	// var data Player

	// if r.Method == "POST" {
	// 	fmt.Println("GET method po raixa")
	// 	fmt.Println(r.URL.Query())
	// 	fmt.Println(r.ParseForm())
	// 	p := Player{
	// 		Id:       1,
	// 		username: "neymarsabin",
	// 		score:    20,
	// 	}
	// 	RespondWithJSON(w, http.StatusOK, p)
	// } else {
	// 	fmt.Println(r.Method + " po raixa")
	// }
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// need to end ) in same line else put comma(,)
func connectionWithPostres() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// db = sql.DB object
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return "Successfully Connected!!"
}
