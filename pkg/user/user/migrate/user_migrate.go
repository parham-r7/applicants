package migrate

import (
	"applicants/models"
	"gorm.io/gorm"
	"log"
)

func DoUserMigrate(db *gorm.DB) {
	err := db.AutoMigrate(models.User{})
	if err != nil{
		log.Println(err)
	}
}
