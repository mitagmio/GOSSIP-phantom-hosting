package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gossip-phantom-hosting/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var account Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	username := account.Username
	password := account.Password

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("Username or password was incorrect")
		utils.Respond(w, nil, err)
		return
	}

	// test
	fmt.Println(username, password)

	utils.Respond(w, account, nil)
}
