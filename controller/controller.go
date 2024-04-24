package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"reddynn/config"
	"reddynn/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to homepage"))

}
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	dbs := config.Dbconnect()
	defer dbs.Close()
	var newuser string
	err = dbs.QueryRow("select username from users where username=?", user.Username).Scan(&newuser)
	switch {
	case err == sql.ErrNoRows:
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "unable to create account", http.StatusInternalServerError)
			return
		}
		_, err = dbs.Query("insert into users(username,password) values(?,?)", user.Username, hashedpassword)
		if err != nil {
			http.Error(w, "server unble to create user", http.StatusInternalServerError)
			fmt.Println(err)
			return

		}
		w.Write([]byte("user has been created"))
	case err != nil:
		http.Error(w, "server db query error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	default:
		http.Error(w, "user already exsited", http.StatusBadRequest)

		return
	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	dbs := config.Dbconnect()
	defer dbs.Close()
	var newpassword string
	err = dbs.QueryRow("select password from users where username=?", user.Username).Scan(&newpassword)
	switch {
	case err != nil:
		http.Error(w, "unauthorized no user with this username", http.StatusUnauthorized)
		return
	default:
		err := bcrypt.CompareHashAndPassword([]byte(newpassword), []byte(user.Password))
		if err != nil {
			http.Error(w, "unathorized password wrong", http.StatusUnauthorized)
			return
		}
		w.Write([]byte("welcome" + user.Username))
	}

}
