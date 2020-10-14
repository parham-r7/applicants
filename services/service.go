package services

import (
	"applicants/database/mysql"
	"applicants/pkg/user/user"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run() {
	var (
		err        error
		db         *gorm.DB
		router     *mux.Router
		apiHandler *mux.Router
		port       = viper.GetString("service.port")
	)
	//	init database
	db, err = mysql.Connect()
	if err != nil {
		panic(err)
	}

	// init handler
	router = mux.NewRouter()
	apiHandler = router.PathPrefix("/api").Subrouter()

	//init services
	user.Init(db, apiHandler)

	log.Fatalln(http.ListenAndServe(port, router))

	//err = http.ListenAndServe(port, router)
	//if err != nil {
	//	log.Println(err)
	//}
}
