package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/clientes", getClientes)
	http.HandleFunc("/cliente", postCliente)

	http.ListenAndServe(":8080", nil)

}

func getClientes(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type","application/json")

	// var clientes = Clientes{
	// 	Cliente{"Matheus"},
	// 	Cliente{"Freire"},
	// }

	var clientes = db()

	json.NewEncoder(w).Encode(clientes)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "senha"
	dbname   = "diodb"
)

func db() User {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var db, _ = sql.Open("postgres", psqlInfo)
	defer db.Close()

	var sqlStatement = `SELECT * FROM users WHERE id=$1;`
	var user User
	var row = db.QueryRow(sqlStatement, 2)

	fmt.Println(row)

	_ = row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)

	fmt.Println(user)

	return user
}

type User struct {
	ID			int `json:"id"`
	Age 		int `json:"age"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
}

func postCliente(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res := Clientes{}
	body, _ := io.ReadAll(r.Body)

	json.Unmarshal(body,&res)

	fmt.Println(res)
	fmt.Println(res[0].Nome)
	fmt.Println(res[1].Nome)
}

type Cliente struct {
	Nome string `json:"name"`

}

type Clientes []Cliente