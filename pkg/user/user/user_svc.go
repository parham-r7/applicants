package user

import (
	"applicants/pkg/user/repository"
	"applicants/pkg/user/transport/http"
	"applicants/pkg/user/usecase"
	migrate2 "applicants/pkg/user/user/migrate"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, router *mux.Router) {
	var (
		userRepo repository.UserRepo
		userUC   usecase.UserUC
		migrate  = viper.GetBool("migrate.user")
	)

	if migrate {
		migrate2.DoUserMigrate(db)
	}

	userRepo = repository.New(db)
	userUC = usecase.New(userRepo)
	http.InitUserHandler(userUC, router)


}
