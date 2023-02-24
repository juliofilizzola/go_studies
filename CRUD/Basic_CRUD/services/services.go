package services

import (
	db2 "Basic_CRUD/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpor"))
	}

	var user user

	if err = json.Unmarshal(request, &user); err != nil {
		w.Write([]byte("Falha ao ler o corpor"))
	}

	fmt.Println(user)

	db, err := db2.Connect()
	if err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}

	statment, err := db.Prepare("insert into user (name, email)values (?, ?)")

	if err != nil {
		w.Write([]byte("erro ao criar o statement!"))
		return
	}

	insert, err := statment.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("erro ao exec o statement!"))
		return
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("erro ao criar o statement!"))
		return
	}
	w.Write([]byte("id"))

	fmt.Println(idInsert)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db2.Connect()
	if err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}

	getUsers, err := db.Query("select * from user")
	if err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}
	defer getUsers.Close()

	var users []user

	for getUsers.Next() {
		var user user
		if err := getUsers.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("error ao connect database"))
			return
		}

		users = append(users, user)

	}
	fmt.Println(users)
	//w.Write(users)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}
}
func GetUser(w http.ResponseWriter, r *http.Request) {}
