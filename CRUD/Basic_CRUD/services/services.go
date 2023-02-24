package services

import (
	db2 "Basic_CRUD/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("parse uint err"))
		return
	}
	fmt.Println(ID)
	db, err := db2.Connect()
	if err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}

	getUser, err := db.Query("select * from user where id = ? ", ID)

	if err != nil {
		w.Write([]byte("erro request db!"))
		return
	}
	var user user
	if getUser.Next() {
		if err := getUser.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("erro request db!"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	fmt.Println(user)
	//w.Write(users)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("error convert response"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("parse uint err"))
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpor"))
	}

	var requestUser user

	if err = json.Unmarshal(request, &requestUser); err != nil {
		w.Write([]byte("Falha ao ler o corpor"))
	}

	fmt.Println(requestUser, "requestUser")
	db, err := db2.Connect()
	if err != nil {
		w.Write([]byte("error ao connect database"))
		return
	}
	defer db.Close()

	upUser, err := db.Prepare("update user set name= ?, email =? where id = ?")

	if err != nil {
		w.Write([]byte("erro ao criar o statement!"))
		return
	}
	defer upUser.Close()

	if _, err := upUser.Exec(requestUser.Name, requestUser.Email, ID); err != nil {
		w.Write([]byte("erro ao exec o statement!"))
		return
	}

	getUser, err := db.Query("select * from user where id = ? ", ID)
	if err != nil {
		w.Write([]byte("erro request db!"))
		return
	}
	defer getUser.Close()
	fmt.Println(getUser, "get US")
	var Upuser user

	if getUser.Next() {
		if err := getUser.Scan(&Upuser.ID, &Upuser.Name, &Upuser.Email); err != nil {
			w.Write([]byte("erro request db!"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	fmt.Println(Upuser, "up user")
	//w.Write(users)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Upuser); err != nil {
		w.Write([]byte("error convert response"))
		return
	}

}
