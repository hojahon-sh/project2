package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/project2/models"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		user = models.User{}
		err  error
	)

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
