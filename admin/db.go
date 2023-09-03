package admin

import (
	"os"
  "fmt"

	"github.com/alexeyismirnov/go-vue-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB) {
	var err error

	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	// Set db log level
	db.Logger = db.Logger.LogMode(logger.Info)

	// Create data table in the database
	err = db.AutoMigrate(models.Message{})
	if err != nil {
		panic(err)
	}

  var cnt int64
	if err := db.Table("messages").Count(&cnt).Error; err != nil {
		panic(err)
	}

	if cnt == 0 {
    print("ADDING MESSAGE")
    message := models.Message{Content: "hello world"}
    db.Create(&message)

    print(message.ID)

  } else {
    print("RETRIEVING MESSAGE")
    var message models.Message
    db.First(&message)

    print(fmt.Sprintf("Message %d - %s\n\n" , message.ID, message.Content))
  }

	return
}
