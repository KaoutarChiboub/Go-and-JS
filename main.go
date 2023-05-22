package main

import (
	"database/sql"  // To connect to sql db
	"encoding/json" // To encode data to json
	"fmt"           // To print out information about servers and connections
	"log"           // To log out errors
	"net/http"      // To handle http requests
	"os"            // To get variables

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
)

type User struct {
	ID    string `json:"id"`
	ICCID string `json:"ICCID"`
	IMSI  string `json:"IMSI"`
	LAI   string `json:"LAI"`
	K     string `json:"K"`
	OEN   string `json:"OEN"`
}

var db *sql.DB

// The following functions are used to handle HTTP requests

func getSimList(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.ICCID, &u.IMSI, &u.LAI, &u.K, &u.OEN); err != nil {
				log.Fatal(err)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		res, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func createSimUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		w.Header().Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(&u)                                                                                                        
		_, err := db.Exec("INSERT INTO users (ICCID, IMSI, LAI, K, OEN) VALUES ($1, $2, $3, $4, $5)", u.ICCID, u.IMSI, u.LAI, u.K, u.OEN)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
		
		//message := "Sim user created successfully"
		//w.Write([]byte(message))
	}
}

func main() {

	//Connection to postgres database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/simlist", getSimList(db)).Methods("GET")
	r.HandleFunc("/sim-login", createSimUser(db)).Methods("POST")

	handler := cors.Default().Handler(r)

	fmt.Printf("Starting the server at port 8008\n")
	log.Fatal(http.ListenAndServe(":8008", handler))

}
