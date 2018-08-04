package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/pacman/model"
	// _ "github.com/lib/pq"
	"log"
	"net/http"
)

// information on database credentials
const (
	host     = "localhost"
	port     = 5432
	user     = "neymarsabin"
	password = "gopacman123"
	dbname   = "postgres"
)

type Pacman struct {
	DB     *pg.DB
	Router *mux.Router
}

// var psqlInfo = connectionWithPostres()

func (p Pacman) Initialize() {
	p.DB = model.Connection()
}

func (p Pacman) InitRoutes() {
	p.Router = mux.NewRouter()
	p.Router.HandleFunc("/", DisplayUserInfo).Methods("GET")
	// passing post method
	p.Router.HandleFunc("/score", p.createUserScore).Methods("POST")
}

func DisplayUserInfo(w http.ResponseWriter, r *http.Request) {
	player := model.Player{
		Id:       1,
		Username: "neymarsabin",
		Score:    20,
	}
	RespondWithJSON(w, http.StatusOK, player)
}

func (p Pacman) createUserScore(w http.ResponseWriter, r *http.Request) {
	var player model.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		fmt.Println("Error occured while decoding payload")
		fmt.Println("Details:", err)
		errorMap := make(map[string]string)
		errorMap["error"] = err.Error()
		RespondWithJSON(w, http.StatusBadRequest, errorMap)
		return
	}

	// fmt.Println(player.Id)
	player.SavePlayer(p.DB)
	RespondWithJSON(w, http.StatusOK, player)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// func RespondWithError(w http.ResponWrite, code int, payload inte)

// need to end ) in same line else put comma(,)
// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s modelname=%s sslmode=disable", host, port, user, password, modelname)

// func connectionWithPostres() *sql.DB {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s modelname=%s sslmode=disable", host, port, user, password, modelname)
// 	model, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected...") // model = sql.DB object
// 	defer model.Close()
// 	err = model.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return model
// }

// func insertScoreOfUserInTable(player Player, psqlInfo *sql.DB) string {
// 	fmt.Println(player)
// 	fmt.Println("Executing query")
// 	// stmt, err := psqlInfo.Prepare("INSERT INTO players(id,name,score) VALUES($1,$2,$3);")
// 	// err = stmt.Exec(1, "neymarsabin", 40)
// 	// stmt, err := psqlInfo.QueryRow("INSERT INTO players (id, username, score) values (player.Id, player.Username, player.Score)")
// 	// _, err = stmt.Exec()
// 	// err := psqlInfo.QueryRow("INSERT INTO players(id,name,score) VALUES($1,$2,$3);", 1, "neymarsabin", 40)
// 	// fmt.Println(res)
// 	fmt.Println("Finished Query")
// 	// fmt.Println(err)
// 	return "hello world"
// }

// func savePlayer(player model.Player) {
// 	fmt.Println("This is an application.")
// 	err := player.SavePlayer()
// 	if err != nil {
// 		fmt.Printf("sano savePlayer ma error %v", err)
// 	}
// }

func main() {
	p := Pacman{}
	p.Initialize() // init configs
	p.InitRoutes() // init routes
	log.Fatal(http.ListenAndServe(":8081", p.Router))
}
