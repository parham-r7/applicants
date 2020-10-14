package http

import (
	"applicants/models"
	"applicants/pkg/user/usecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type userHandler struct {
	UC usecase.UserUC
}

func InitUserHandler(userUC usecase.UserUC, router *mux.Router)  {
	var (
		handler = &userHandler{
			UC: userUC,
		}
		apiHandler *mux.Router
		//err error
	)

	apiHandler = router.PathPrefix("/users").Subrouter()
	apiHandler.HandleFunc("", handler.create).Methods("POST")

}

func (h *userHandler)create( w http.ResponseWriter, r *http.Request)  {
	var (
		err error
		user = &models.User{}
	)

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println(err)
	}

	user, err = h.UC.Create(user)
	if err != nil {
		log.Println(err)
	}

	resp := &models.Response{
		Code: http.StatusOK,
		Data: user,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}