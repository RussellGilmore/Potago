package seed

import (
	"log"

	"github.com/RussellGilmore/potago/api/models"
	"github.com/jinzhu/gorm"
)

var potatoes = []models.Potato{
	models.Potato{
		Name:        "Red Potato",
		Description: "A good red potato.",
	},
	models.Potato{
		Name:        "Yellow Potato",
		Description: "A good yellow potato.",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Potato{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Potato{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range potatoes {
		err = db.Debug().Model(&models.Potato{}).Create(&potatoes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed potatoes table: %v", err)
		}
	}
}
